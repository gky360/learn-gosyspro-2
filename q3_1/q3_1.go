package main

import (
	"io"
	"os"
)

func main() {
	inFile, err := os.Open("./q3_1/q3_1.go")
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	outFile, err := os.Create("./tmp/q3_1.txt")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, inFile)
	if err != nil {
		panic(err)
	}
}
