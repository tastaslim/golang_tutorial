package dataTypes

import (
	"fmt"
)

func TestDataTypes() {
	// var num uint64 = 20 // unsigned Integer
	// fmt.Println(num)

	var num int = 30 // signed integer
	fmt.Println(num)

	var name string = "Dice" // string
	var msg string = fmt.Sprintf("Hi there my name is %v", name)
	fmt.Println(msg)

	var salary float32 = 1920.20 // float (float32, float64)
	fmt.Println(salary)

	var isMale bool = true // bool
	fmt.Println(isMale)

	var complexType complex64 = 10 + 12i
	fmt.Println(complexType)

	var complexTypeOtherWay complex64 = complex(10, 2)
	fmt.Println(complexTypeOtherWay)

	// summation
	summ := complexType + complexTypeOtherWay
	fmt.Println(summ)

	// Get real and Imaginary Part

	realPart := real(complexType)
	imaginaryPart := imag(complexType)

	fmt.Println(realPart, imaginaryPart)

	/*
		Type Casting in Golang
	*/

	var toBeCasted float32 = 1235.476
	castedNumber := int(toBeCasted)
	fmt.Println(castedNumber)

}
