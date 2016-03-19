// Copyright 2016 fluter

package paste

import "testing"

type entry struct {
	url string
	ok  bool
}

var entries []entry = []entry{
	{"http://codepad.org/abcd", false},
	{"http://codepad.org/d4TzyKaI", true},
	{"http://bpaste.net/show/abc", false},
	{"https://bpaste.net/show/0313375aa3c8", true},
	{"http://pastebin.com/abc", false},
	{"http://pastebin.com/8KACZxp7", true},
	{"http://sprunge.us/abc", true},
	{"http://sprunge.us/BBai", true},
	{"http://ideone.com/abc", true},
	{"http://ideone.com/y2FsS8", true},
	{"http://pastie.org/abc", false},
	{"http://pastie.org/10764925", true},
	{"http://dpaste.com/abc", false},
	{"http://dpaste.com/2CQFK52", true},
	{"http://foo.bar/baz", false},
}

func TestGet(t *testing.T) {
	for _, e := range entries {
		t.Log("Getting", e.url)
		ret, err := Get(e.url)
		if err != nil {
			t.Log(err)
		} else {
			t.Log(ret)
		}
		if (err == nil) == e.ok {
			t.Log("OK")
		} else {
			t.Fail()
			t.Log("FAIL")
		}
	}
}

func TestPut(t *testing.T) {
	text := "This is a test message"
	id, err := Paste(text)
	if err != nil {
		t.Error(err)
	}
	t.Log(id)
}
