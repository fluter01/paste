package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/fluter01/paste/sprunge"
)

func usage(prog string) {
	fmt.Printf("Usage:    %s [-h|--help] <file>\n\n"+
		"             -h|--help - Show this help message.\n"+
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

	res, err = sprunge.Put(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("URL: ", res)
}

func main() {
	var help bool

	flag.BoolVar(&help, "h", false, "Show help message")
	flag.BoolVar(&help, "help", false, "Show help message")
	flag.Parse()

	if help {
		usage(os.Args[0])
	}

	for i := 0; i < flag.NArg(); i++ {
		file := flag.Arg(i)
		processFile(file)
	}
	return
}
