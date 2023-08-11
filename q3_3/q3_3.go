package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("./tmp/q3_3.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	writer, err := zipWriter.Create("a.txt")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(writer, strings.NewReader("a.text contents"))
	if err != nil {
		panic(err)
	}
}
