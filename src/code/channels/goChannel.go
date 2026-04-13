package main

import (
	"fmt"
	"time"
)

func processNumber(num chan int) {
	fmt.Printf("Processing number %v", <-num)
}
func main() {
	numChannel := make(chan int)
	go processNumber(numChannel)
	numChannel <- 5
	time.Sleep(time.Second * 2)
}
