package goroutines

import (
	"fmt"
)

func PrintNumbers(num int) {
	fmt.Println(num)

}

// // The below code will exit without printing numbers or partial printing numbers because the main function will exit before the goroutines can print the numbers. We will learn about wait groups to handle this.
// // func main() {
// // 	for i := 0; i <= 10; i++ {
// // 		go PrintNumbers(i)
// // 	}
// // }
