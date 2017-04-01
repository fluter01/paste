# paste
[![GoDoc](https://godoc.org/github.com/fluter01/paste?status.svg)](https://godoc.org/github.com/fluter01/paste)
[![Build Status](https://travis-ci.org/fluter01/paste.svg?branch=master)](https://travis-ci.org/fluter01/paste)
[![codecov.io](https://codecov.io/github/fluter01/paste/coverage.svg?branch=master)](https://codecov.io/github/fluter01/paste?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/fluter01/paste)](https://goreportcard.com/report/github.com/fluter01/paste)
[![Gitter](https://badges.gitter.im/fluter01/paste.svg)](https://gitter.im/fluter01/paste?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)

## A tool and API downloading from and send paste to online pastebin services.

Currently support getting paste from following pastebins:

* bpaste.net
* codepad.org
* ideone.com
* pastebin.com
* pastie.org
* sprunge.us
* privatepaste.com
* paste.debian.net
* paste.fedoraproject.org
* ptpb.pw
* paste.pr0.tips
* vp.dav1d.de
* hastebin.com
* lpaste.net
* fpaste.org
* ghostbin.com
* dpaste.de
* codeviewer.org
* paste.ee
* paste.linuxassist.net
* paste.linux.chat
* paste.pound-python.org
* pastebin.geany.org
* paste.kde.org
* paste.eientei.org
* www.heypasteit.com
* paste.ubuntu.org.cn
* pastebin.ca
* paste.lugons.org
* play.golang.org
* glot.io
* vpaste.net
* ix.io

Send paste to:
* sprunge.us

## Usage:

### download paste

```bash
gopaste <paste url>
```

### send paste

Send a file:
```bash
gopaste -p foo.txt
```
Send a command's output:
```bash
cmd | gopaste -p -
```

## Why

I'm an IRC user that works on terminal, when someone post his code snippet in the channel, I hate to lunch the browser to view that code. So I decided to make this tool. It will help you get the raw content of the snippets in your terminal, no browser, no GUI needed!
The project is partly inspired by [fstd/cwarn](https://github.com/fstd/cwarn), which scrapes paste links and tries to compile them.

