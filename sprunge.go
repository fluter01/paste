package sprunge

import (
    "errors"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "regexp"
)

const SPRUNGE_URL = "http://sprunge.us"
var re = regexp.MustCompile("http://sprunge.us/([[:alnum:]]+)")


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
    var resp *http.Response
    var body []byte
    var url string

    url = fmt.Sprintf("%s/%s", SPRUNGE_URL, id)

    resp, err = http.Get(url)
    if err != nil {
        fmt.Println("error while sending data to sprunge")
        fmt.Println(err);
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
    return string(body), nil
}

func Put(data string) (string, error) {
    var err error
    var resp *http.Response
    var formData url.Values
    var body []byte

    formData = url.Values {"sprunge" : {data}}

    resp, err = http.PostForm(SPRUNGE_URL, formData)
    if err != nil {
        fmt.Println("error while sending data to sprunge")
        fmt.Println(err);
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
    return string(body), nil
}
