package goroutines

import (
	"fmt"
	"sync"
)

func PrintFibonacci(numberToPrint int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	fmt.Println(numberToPrint)
}

func TestWaitGroups() {
	var waitGroup sync.WaitGroup
	num := 1
	for number := range num {
		waitGroup.Add(1)
		go PrintFibonacci(number, &waitGroup)
	}
	waitGroup.Wait()
}
