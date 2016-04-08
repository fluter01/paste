// Copyright 2016 fluter

package paste

import "testing"

type entry struct {
	url string
	ok  bool
}

var entries = []entry{
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
	{"http://paste.debian.net/427049/", true},
	{"http://paste.debian.net/abc", false},
	{"http://paste.fedoraproject.org/351674/", true},
	{"http://paste.fedoraproject.org/false/", false},
	{"http://paste.fedoraproject.org/351675/16030146/", true},
	{"https://ptpb.pw/WMa-", true},
	{"https://ptpb.pw/false", false},
	{"http://paste.pr0.tips/Ky", true},
	{"http://paste.pr0.tips/false", true},
	{"http://vp.dav1d.de/Vvn4U", true},
	{"http://vp.dav1d.de/false", false},
	{"http://hastebin.com/ilanaqadel.vbs", true},
	{"http://lpaste.net/159275", true},
	{"http://fpaste.org/351700/", true},
	{"https://ghostbin.com/paste/tzh3m", true},
	{"https://dpaste.de/DhN3", true},
	{"http://paste.ee/p/yIw9B", true},
}

func TestGet(t *testing.T) {
	for _, e := range entries[15:] {
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
