package main

import (
	"fmt"
)

func PrintNumbers(num int) {
	fmt.Println(num)

}

/*
The below code will exit without printing numbers. What really happens it, the main function(say main thread) is also running inside go routine
func main() {
	for i := 0; i <= 10; i++ {
		go PrintNumbers(i)
	}
}
*/
