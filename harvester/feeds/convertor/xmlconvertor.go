package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	xj "github.com/basgys/goxml2json"
)

func main() {
	// First element in os.Args is always the program name,
	// So we need at least 2 arguments to have a file name argument.
	if len(os.Args) < 2 {
		os.Stderr.Write([]byte("Missing parameter, provide file name!"))
		return
	}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		os.Stderr.Write([]byte("Can't read file"))
		return
	}

	xml := bytes.NewReader(data)
	jsbuf, err := xj.Convert(xml)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(jsbuf)
}
