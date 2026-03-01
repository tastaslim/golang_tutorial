package loops

import (
	"fmt"

	"git.druva.org/cloudapps/shareddrive-node/src/pkg/logger"
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

func RunMiscellenious() {
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
