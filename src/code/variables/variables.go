package variables

import (
	"fmt"
)

const (
	a = 20
	b = 30
	c = 40
)

func VariablePractice() {
	var number1 int = 32
	var number2 int = 34
	fmt.Println(number1 + number2)
	// fmt.Print("Hi")
	// fmt.Printf("Hi")

	a := 32
	b := 43
	fmt.Println(a + b)

	// c:=  ==> Wrong, it is must to assing it some value while declaring variable in this way

	var c int
	c = 4
	fmt.Println(c)

	/* Declaring multiple variables in one line*/

	// var c int, var tas int = 8, var dinky string = "Dinky", var sam bool = true, var ankit float32 = 8.4 ==> Incorrect because We can not define variables of multiple type like this

	var (
		dinky string = "Dinky"
		tas   int    = 8
		sam   bool   = true
	)

	fmt.Println(dinky, sam, tas)

	// As a best practice, we will always define variables of one type in one line

	// var a1, b1 int = 1, 2
	a1, b1 := 1, 1 // Best practice if variable scope is within method
	fmt.Println(a1, b1)

}

// Variable Shadowing

func TestVariableShadowing() {
	var x, y int = 10, 20
	if true {
		var x int = 5 // Variable Shadowing
		fmt.Println(x)
		y = 25
		fmt.Println(y)
	}
	fmt.Println(x)
	fmt.Println(y)
}

func TestConstant() {
	const num int = 20 // const num = 20 ==> Both are correct
	fmt.Println(num)
	var name string = "James Bond"
	fmt.Println(name)
	fmt.Println(a, b, c)
}

// name := "James Bond" ==> Wrong will give error
// var name string = "James Bond"
// fmt.Println(name)
