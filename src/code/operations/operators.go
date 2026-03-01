package operations

import (
	"fmt"
)

func RunOperations() {
	// Arithmetic Operators
	var a, b int = 10, 3
	fmt.Println("Arithmetic Operators:")
	fmt.Println("Addition:", a+b)       // 13
	fmt.Println("Subtraction:", a-b)    // 7
	fmt.Println("Multiplication:", a*b) // 30
	fmt.Println("Division:", a/b)       // 3
	fmt.Println("Modulus:", a%b)        // 1

	// Relational Operators
	fmt.Println("\nRelational Operators:")
	fmt.Println("Equal:", a == b)            // false
	fmt.Println("Not Equal:", a != b)        // true
	fmt.Println("Greater:", a > b)           // true
	fmt.Println("Less:", a < b)              // false
	fmt.Println("Greater or Equal:", a >= b) // true
	fmt.Println("Less or Equal:", a <= b)    // false

	// Logical Operators
	fmt.Println("\nLogical Operators:")
	fmt.Println("AND:", a > b && b > 0) // true
	fmt.Println("OR:", a < b || b > 0)  // true
	fmt.Println("NOT:", !(a == b))      // true

	// Bitwise Operators
	fmt.Println("\nBitwise Operators:")
	fmt.Println("AND:", a&b)          // 2
	fmt.Println("OR:", a|b)           // 11
	fmt.Println("XOR:", a^b)          // 9
	fmt.Println("Left Shift:", a<<1)  // 20
	fmt.Println("Right Shift:", a>>1) // 5

	// Assignment Operators
	fmt.Println("\nAssignment Operators:")
	c := a
	c += 5
	fmt.Println("+=:", c) // 15
	c -= 3
	fmt.Println("-=:", c) // 12
	c *= 2
	fmt.Println("*=:", c) // 24
	c /= 4
	fmt.Println("/=:", c) // 6
}
