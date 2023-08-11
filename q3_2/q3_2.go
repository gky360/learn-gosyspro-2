package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	file, err := os.Create("./tmp/q3_2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.CopyN(file, rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
}
