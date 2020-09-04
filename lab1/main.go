package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func prettyPrintf(format string, args ...interface{}) {
	data := fmt.Sprintf(format, args...)
	line := "\n" + strings.Repeat("=", len(data)) + "\n"
	fmt.Printf("%s%s%s\n", line, data, line)
}

func main() {
	var testdir string

	flag.StringVar(&testdir, "t", "tests", "папка с тестами")
	flag.Parse()

	for _, labName := range flag.Args() {
		var f func(io.Reader)
		// TODO : func (r io.Reader, feature bool)
		switch labName {
		case "lab1":
			f = lab1
		case "lab2":
			f = lab2
		case "lab3":
			f = lab3
		case "lab4":
			f = lab4
		case "lab5":
			f = lab5
		default:
			log.Fatal("Нет такой лабораторной: ", labName)
		}

		path := filepath.Join(testdir, labName)
		files, err := ioutil.ReadDir(path)
		if err != nil {
			log.Print("Error opening dir: ", err)
		}

		for _, file := range files {
			filePath := filepath.Join(path, file.Name())
			prettyPrintf("%s>  %s", labName, filePath)

			opened, err := os.Open(filePath)
			if err != nil {
				log.Fatal("Error opening file: ", err)
			}
			f(opened)
		}
	}
}
