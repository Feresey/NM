package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func prettyPrint(format string, args ...interface{}) {
	data := fmt.Sprintf(format, args...)

	fmt.Print("\n")
	for i := 0; i < len(data); i++ {
		fmt.Print("=")
	}
	fmt.Print("\n")
	fmt.Print(data)
	fmt.Print("\n")
	for i := 0; i < len(data); i++ {
		fmt.Print("=")
	}
	fmt.Print("\n\n")
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
		case "lab5":
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
			prettyPrint("%s>  %s", labName, filePath)

			opened, err := os.Open(filePath)
			if err != nil {
				log.Print("Error opening file: ", err)
				continue
			}
			f(opened)
		}
	}
}
