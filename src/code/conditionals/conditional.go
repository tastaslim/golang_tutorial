package conditionals

import (
	"fmt"
)

func addition(a int, b int) int { // Private Method
	return a + b
}

func ArithmeticAddition(a int, b int) { // Public Method
	summ := addition(a, b)
	fmt.Println(summ)
}

func ConditionalOperations() {
	var a int = 10 // a:= 10
	b := 20        // var b int = 20
	if a > b {
		var logMessage string = fmt.Sprintf("%d is greater than %d", a, b)
		fmt.Println(logMessage)
	} else if a < b {
		var logMessage string = fmt.Sprintf("%d is lesser than %d", a, b)
		fmt.Println(logMessage)
	} else {
		var logMessage string = fmt.Sprintf("%d is equal to %d", a, b)
		fmt.Println(logMessage)
	}
}

func SwitchStatement() {
	a := 4
	switch a {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	}
}
