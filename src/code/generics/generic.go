package generics

import "fmt"

/*
Both are same
*/
func AddElements[T any](items []T) T { // func AddElements[T interface{}](items []T) T
	var finalResult T
	for i := range items {
		fmt.Println(items[i])
	}
	return finalResult
}

func AddElementsUsingInterfaceInsteadOfAny[T interface{}](items []T) T { // func AddElements[T any](items []T) T
	var finalResult T
	for i := range items {
		fmt.Println(items[i])
	}
	return finalResult
}

/*
Say I want to only allow int and string types
*/

func AddElementOnlyStringAndInt[T int | string](items []T) []T {
	fmt.Println(items)
	return items
}

// Generics in struct
/*
type Stack struct {
	elements []int
}

func main() {
	myStack := Stack{
		elements: []int{1, 2, 3},
	}
	fmt.Println(myStack)
}
*/

type Stack[T any] struct {
	elements []T
}

func main() {
	myStack := Stack[int]{
		elements: []int{1, 2, 3, 4},
	}
	fmt.Println(myStack)
}
