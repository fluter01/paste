package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/fluter01/paste/bpaste"
	"github.com/fluter01/paste/codepad"
	"github.com/fluter01/paste/ideone"
	"github.com/fluter01/paste/pastebin"
	"github.com/fluter01/paste/sprunge"
)

func usage(prog string) {
	fmt.Printf("Usage:    %s [-h|--help] -s <service> -g|-p <file>\n\n"+
		"             -h|--help - Show this help message.\n"+
		"             -s|--service - Select paste service(Default sprunge).\n"+
		"             -g|--get - Get the paste.\n"+
		"             -p|--put - Send the paste.\n"+
		"             file      - The file that will be sent to paste.\n",
		filepath.Base(prog))
	os.Exit(1)
}

func readFromStdin() (string, error) {
	var data []byte
	var err error

	data, err = ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("failed to read from stdin")
		fmt.Println(err)
		return "", err
	}
	return string(data), nil
}

func readFromFile(filename string) (string, error) {
	var data []byte
	var err error

	data, err = ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("failed to read from file " + filename)
		fmt.Println(err)
		return "", err
	}
	return string(data), nil
}

func processFile(filename string) {
	var err error
	var data string
	var res string

	if filename == "-" {
		data, err = readFromStdin()
	} else {
		data, err = readFromFile(filename)
	}

	if err != nil {
		return
	}

	res, err = putter(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("URL:", res)
}

func processID(id string) {
	var err error
	var res string

	res, err = getter(id)
	if err != nil {
		return
	}

	fmt.Println(res)
}

var (
	getter func(string) (string, error)
	putter func(string) (string, error)
)

func main() {
	var help bool
	var service string
	var get bool
	var put bool

	flag.BoolVar(&help, "h", false, "Show help message")
	flag.BoolVar(&help, "help", false, "Show help message")
	flag.StringVar(&service, "s", "sprunge", "Paste service")
	flag.StringVar(&service, "service", "sprunge", "Paste service")
	flag.BoolVar(&get, "g", false, "Get the paste")
	flag.BoolVar(&get, "get", false, "Get the paste")
	flag.BoolVar(&put, "p", false, "Send the paste")
	flag.BoolVar(&put, "put", false, "Send the paste")
	flag.Parse()

	if help || (get && put) || !(get || put) {
		usage(os.Args[0])
	}

	switch service[0] {
	default:
		fallthrough
	case 's':
		getter = sprunge.Get
		putter = sprunge.Put
	case 'p':
		getter = pastebin.Get
		putter = pastebin.Put
	case 'c':
		getter = codepad.Get
		putter = codepad.Put
	case 'i':
		getter = ideone.Get
		putter = ideone.Put
	case 'b':
		getter = bpaste.Get
		putter = bpaste.Put
	}

	if put {
		for i := 0; i < flag.NArg(); i++ {
			file := flag.Arg(i)
			processFile(file)
		}
	} else {
		for i := 0; i < flag.NArg(); i++ {
			id := flag.Arg(i)
			processID(id)
		}
	}
	return
}
