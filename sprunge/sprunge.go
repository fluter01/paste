package sprunge

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

const SPRUNGE_URL = "http://sprunge.us"

var re = regexp.MustCompile("http://sprunge.us/([[:alnum:]]+)")

// Extract sprunge ID from the given URL.
func GetID(url string) (string, error) {
	var match []string

	match = re.FindStringSubmatch(url)
	if len(match) != 2 {
		return "", errors.New("invalid sprunge url")
	}
	return match[1], nil
}

// Get the data from the paste ID.
// This returns the entire data as a string.
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

// Get the reader for the paste ID.
// The caller can then read from this reader, when reading is done,
// the caller need to close the reader.
func GetReader(id string) (io.Reader, error) {
	var err error
	var resp *http.Response
	var url string

	url = fmt.Sprintf("%s/%s", SPRUNGE_URL, id)

	resp, err = http.Get(url)
	if err != nil {
		fmt.Println("error while get paste from sprunge:", err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		fmt.Println("server returned an error:", resp.Status)
		return nil, errors.New(resp.Status)
	}

	return resp.Body, nil
}

// Send the data string to paste server.
// Returns the ID of this paste if succeeds.
func Put(data string) (string, error) {
	var err error
	var resp *http.Response
	var formData url.Values
	var body []byte

	formData = url.Values{"sprunge": {data}}

	resp, err = http.PostForm(SPRUNGE_URL, formData)
	if err != nil {
		fmt.Println("error while sending data to sprunge")
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("server returned an error:")
		fmt.Println(resp.Status)
		return "", errors.New(resp.Status)
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read server response error")
		fmt.Println(err)
		return "", err
	}
	return string(bytes.TrimSpace(body)), nil
}
