package loops

import (
	"fmt"

	"github.com/tastaslim/golang_tutorial/src/pkg/logger"
)

func LoopOperations() {
	logger.Info("Info")
	count := 10
	for i := range count { // modern way
		fmt.Println(i)
	}
	/*
		// Legacy
		for i := 1; i < count; i++ {
			fmt.Println(i)
		}

		for i := 1; i < count; i+=1 {
			fmt.Println(i)
		}
	*/
}

func RunMiscellaneous() {
	books := 30
	for book := range books {
		if book > 10 {
			fmt.Println(book)
		} else if book == 10 {
			break
		} else {
			fmt.Println(book)
		}
	}
}

func EvenNumbers(lastNumber int) {
	for number := range lastNumber + 1 { // Inclusive of last number
		if number%2 == 0 {
			fmt.Println(number)
		}
	}
}

func BreakOnCondition(target int) {
	number := 0
	for {
		if number == target {
			break
		}
		fmt.Println(number)
		number += 1
	}
}

func ReturnOnCondition(targetName string) string {
	names := []string{"Tas", "Mom", "Dad", "Dinky", "Ash"}
	for i := range names { // for i := range len(names) both are correct
		if names[i] == targetName {
			return "YES"
		}
	}
	return "NO"
}
