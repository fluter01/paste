package pastebin

import (
	"errors"
	"fmt"
	"io/ioutil"
	"io"
	"net/http"
	"regexp"
)

const PASTEBIN_URL = "http://pastebin.com"

var re = regexp.MustCompile(PASTEBIN_URL + "/([[:alnum:]]+)")


func GetID(url string) (string, error) {
	var match []string

	match = re.FindStringSubmatch(url)
	if len(match) != 2 {
		return "", errors.New("invalid sprunge url")
	}
	return match[1], nil
}

func Get(id string) (string, error) {
	var err error
	var reader io.Reader
	var body []byte

	reader, err = GetReader(id)
	if err != nil {
		return "", err
	}

	body, err = ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println("read server response error")
		fmt.Println(err)
		return "", err
	}
	return string(body), nil
}

func GetReader(id string) (io.Reader, error) {
	var err error
	var resp *http.Response
	var url string

	url = fmt.Sprintf("%s/raw/%s", PASTEBIN_URL, id)

	resp, err = http.Get(url)
	if err != nil {
		fmt.Println("error while get paste from pastebin:", err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		fmt.Println("server returned an error:", resp.Status)
		return nil, errors.New(resp.Status)
	}

	return resp.Body, nil
}

func Put(data string) (string, error) {
	return "", errors.New("Not implemented")
}
