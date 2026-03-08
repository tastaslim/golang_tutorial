package functions

import (
	"fmt"
)

func TupleSum(a int, b int) (int, int) {
	fmt.Println()
	return a + b, a - b
}

// func modify(x int) {
//     x = 100  // only modifies local copy
// }

// num := 5
// modify(num)
// fmt.Println(num) // 5

// AnonymousMethod Method
var AnonymousFunc = func() string {
	return "Anonymous function"
}
