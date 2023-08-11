package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.After(10 * time.Second)
	fmt.Println("Started timer")
	<-t
	fmt.Println("Stopped timer")
}
