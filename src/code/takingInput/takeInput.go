package takinginput

import (
	// "bufio"
	"fmt"
	// "os"
)

func TakeInput() string {
	// var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	// fmt.Print("Enter input: ")
	// scanner.Scan()
	// var text string = scanner.Text()
	// return text
	var text string
	fmt.Print("Enter input: ")
	fmt.Scanln(&text) // fmt.Scanf("%s", &text) // f.Scan(&text)
	return text
}
