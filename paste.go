// Copyright 2016 fluter

package paste

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	urlparser "net/url"
	"regexp"
)

// The error to indicate getting paste from the site is
// not yet supported. Create an issue on github.com :)
var NotSupported = errors.New("Not supported")

type replace struct {
	from, to string
}

// list of pastebins and substitutions for getting raw content
var pastebins map[string]interface{} = map[string]interface{}{
	"bpaste.net":   replace{"show", "raw"},
	"codepad.org":  "%s/raw.c",
	"dpaste.com":   "%s.txt",
	"ideone.com":   "/plain%s",
	"pastebin.com": "/raw%s",
	"pastie.org":   "/pastes%s/download",
	"sprunge.us":   "%s",
}

// Get the raw content of the paste given in url
// Returns the raw text and errors if any
func Get(url string) (string, error) {
	reader, err := GetReader(url)
	if err != nil {
		return "", err
	}

	if rc, ok := reader.(io.ReadCloser); ok {
		defer rc.Close()
	}
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Get a io.Reader for reading the paste
func GetReader(url string) (io.Reader, error) {
	var err error
	var u *urlparser.URL
	var newpath string
	var newurl string

	u, err = urlparser.Parse(url)
	if err != nil {
		return nil, err
	}

	sub, ok := pastebins[u.Host]
	if !ok {
		return nil, NotSupported
	}

	switch sub.(type) {
	case string:
		rep := sub.(string)
		newpath = fmt.Sprintf(rep, u.Path)
	case replace:
		rep := sub.(replace)
		re := regexp.MustCompile(rep.from)
		newpath = re.ReplaceAllString(u.Path, rep.to)
	default:
		return nil, NotSupported
	}
	u.Path = newpath
	newurl = u.String()

	var rsp *http.Response

	rsp, err = http.Get(newurl)
	if err != nil {
		return nil, err
	}

	if rsp.StatusCode != 200 {
		defer rsp.Body.Close()
		return nil, errors.New(rsp.Status)
	}

	return rsp.Body, nil
}

const Sprunge = "http://sprunge.us"

// Send text to paste.
// Returns the paste URL or error if any.
func Paste(text string) (string, error) {
	var err error
	var resp *http.Response
	var body []byte

	formData := urlparser.Values{"sprunge": {text}}

	resp, err = http.PostForm(Sprunge, formData)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New(resp.Status)
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bytes.TrimSpace(body)), nil
}
