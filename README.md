# paste
[![GoDoc](https://godoc.org/github.com/fluter01/paste?status.svg)](https://godoc.org/github.com/fluter01/paste)
[![Build Status](https://travis-ci.org/fluter01/paste.svg?branch=master)](https://travis-ci.org/fluter01/paste)
[![Go Report Card](https://goreportcard.com/badge/github.com/fluter01/paste)](https://goreportcard.com/report/github.com/fluter01/paste)

## A tool and API downloading from and send paste to online pastebin services.

Currently support getting paste from following pastebins:

* bpaste.net
* codepad.org
* ideone.com
* pastebin.com
* pastie.org
* sprunge.us

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
