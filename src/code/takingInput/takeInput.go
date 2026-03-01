package takinginput

import (
	"bufio"
	"fmt"
	"os"
)

func TakeInput() string {
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	fmt.Print("Enter input: ")
	scanner.Scan()
	var text string = scanner.Text()
	return text
}
