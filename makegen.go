package main

import "path/filepath"

import "os"

import "log"

import "fmt"

import "strings"

import "bufio"

func main() {
	var files []string
	root := "."

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	createOrDelete(files)

}

func createOrDelete(files []string) {
	if _, err := os.Stat("build/gen"); !os.IsNotExist(err) {
		os.RemoveAll("build/gen/")
	}

	os.MkdirAll("build/gen", os.ModePerm)

	f, err := os.Create("build/gen/sources.mk")
	if err != nil {
		log.Fatal(err)
	}

	w := bufio.NewWriter(f)

	w.WriteString(fmt.Sprintf("SRC = \\\n"))

	for _, files := range files {
		if strings.HasSuffix(files, ".c") || strings.HasSuffix(files, ".s") {
			w.WriteString(fmt.Sprintf("\t../%s \\\n", files))
		}
	}

	w.Flush()

}
