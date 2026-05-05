package main

import (
	"fmt"
)

func UnBufferedPractice() {
	unBufferedChannel := make(chan string) // UnBuffered Channel
	go func() {
		unBufferedChannel <- "My name is Tashit Shah"
	}()
	tas := <-unBufferedChannel        // Blocking call
	fmt.Printf("Received: %s\n", tas) // This will never reach and result into deadlock as above call is blocking
}

func BufferedPractice() {
	bufferedChannel := make(chan string, 3) // Buffered Channel
	bufferedChannel <- "My name is Tashit Shah"
	tas := <-bufferedChannel // Non Blocking call
	fmt.Printf("Received: %s\n", tas)
}

func main() {
	UnBufferedPractice()
	BufferedPractice()
}
