package main

import (
	"encoding/csv"
	"os"
)

func main() {
	w1 := csv.NewWriter(os.Stdout)
	w1.Write([]string{"hoge", "fuga"})
	w1.Flush()

	file, err := os.Create("./tmp/q2_2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w2 := csv.NewWriter(file)
	w2.Write([]string{"hoge", "fuga"})
	w2.Flush()
}
