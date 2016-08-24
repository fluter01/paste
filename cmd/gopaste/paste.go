// Copyright 2016 fluter

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/fluter01/paste"
)

func usage(prog string) {
	fmt.Printf("Usage:    %s [-h|--help] -s <service> -g|-p <file>\n\n"+
		"             -h|--help - Show this help message.\n"+
		"             -s|--service - Select paste service(Default sprunge).\n"+
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

var putter func(string) (string, error)

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

	res, err = paste.Paste(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("URL:", res)
}

func main() {
	var help bool
	var service string
	var put bool

	flag.BoolVar(&help, "h", false, "Show help message")
	flag.BoolVar(&help, "help", false, "Show help message")
	flag.StringVar(&service, "s", "sprunge", "Paste service")
	flag.StringVar(&service, "service", "sprunge", "Paste service")
	flag.BoolVar(&put, "p", false, "Send the paste")
	flag.BoolVar(&put, "put", false, "Send the paste")
	flag.Parse()

	if help {
		usage(os.Args[0])
	}

	if !put && flag.NArg() == 0 {
		usage(os.Args[0])
	}

	if put {
		for i := 0; i < flag.NArg(); i++ {
			file := flag.Arg(i)
			processFile(file)
		}
	} else {
		var sep bool = flag.NArg() > 1
		for i := 0; i < flag.NArg(); i++ {
			url := flag.Arg(i)
			if sep {
				fmt.Printf("%s:\n", url)
			}
			text, err := paste.Get(url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
				continue
			}
			fmt.Println(text)
		}
	}
	return
}
