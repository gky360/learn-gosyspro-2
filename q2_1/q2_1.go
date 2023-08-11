package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("./tmp/q2_1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(file, "%d, %s, %f\n", 1, "hoge", 2.0)
}
