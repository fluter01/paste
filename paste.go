// Copyright 2016 fluter

// Package paste provides APIs to download and send paste snippets
// from/to online pastebin services.
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

const sprunge = "http://sprunge.us"

// The error to indicate getting paste from the site is
// not yet supported. Create an issue on github.com :)
var ErrNotSupported = errors.New("Not supported")

type replace struct {
	from, to string
}

// list of pastebins and substitutions for getting raw content
var pastebins = map[string]interface{}{
	"bpaste.net":              replace{"show", "raw"},
	"codepad.org":             "%s/raw.c",
	"dpaste.com":              "%s.txt",
	"ideone.com":              "/plain%s",
	"pastebin.com":            "/raw%s",
	"pastie.org":              "/pastes%s/download",
	"sprunge.us":              "%s",
	"privatepaste.com":        "/download%s",
	"paste.debian.net":        "/plain%s",
	"paste.fedoraproject.org": "%s/raw",
	"ptpb.pw":                 "%s",
	"paste.pr0.tips":          "%s",
	"vp.dav1d.de":             "%s",
	"hastebin.com":            replace{"/([^.]*)\\.([a-zA-Z0-9]+)", "/raw/$1"},
	"lpaste.net":              "/raw%s",
	"fpaste.org":              "%s/raw/",
	"ghostbin.com":            "%s/raw",
	"dpaste.de":               "%s/raw",
	"codeviewer.org":          replace{"view", "download"},
	"paste.ee":                replace{"/p/", "/r/"},
	"paste.linuxassist.net":   replace{"view", "view/raw"},
	"paste.linux.chat":        replace{"view", "view/raw"},
	"paste.pound-python.org":  replace{"show", "raw"},
	"pastebin.geany.org":      "%s/raw",
	"paste.kde.org":           "%s/raw/raw",
	"paste.eientei.org":       replace{"show", "raw"},
	"www.heypasteit.com":      replace{"clip", "download"},
	"paste.ubuntu.org.cn":     replace{"/([a-zA-Z0-9]+)", "d$1"},
	"pastebin.ca":             "/raw%s",
	"paste.lugons.org":        replace{"show", "raw"},
	"play.golang.org":         "%s.go",
	"glot.io":                 "%s/raw",
	"vpaste.net":              "%s",
	//	"pastebin.mozilla.org":    replace{"/(.*)", "/?dl=$1"},
}

// Get download the raw content of the paste given in url.
// Returns the raw text and errors if any.
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

// GetReader returns a io.Reader for reading the paste.
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
		return nil, ErrNotSupported
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
		return nil, ErrNotSupported
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

// Paste sends text to paste.
// Returns the paste URL or error if any.
func Paste(text string) (string, error) {
	var err error
	var resp *http.Response
	var body []byte

	formData := urlparser.Values{"sprunge": {text}}

	resp, err = http.PostForm(sprunge, formData)
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
