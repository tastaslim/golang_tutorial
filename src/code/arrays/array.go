package arrays

import (
	"fmt"
	"slices"
	"strings"

	"git.druva.org/cloudapps/shareddrive-node/src/pkg/logger"
)

func ArrayOperations() {
	var array1 = [3]int{1, 2, 3}
	fmt.Println(array1)
	var arr1 = []int{1, 2, 3}
	arr2 := []int{4, 5, 6, 7, 8}
	fmt.Println(arr1) // {1,2,3}
	fmt.Println(arr2) // {4,5,6,7,8}
	arr1[0] = 3       // {3,2,3}
	fmt.Println(arr1)
	arr1 = append(arr1, 90) // {3,2,3,90}
	slices.Reverse(arr1)    // {90,3,2,3}
	var length int = len(arr1)
	fmt.Println(length)

	logger.Debug("Hi There")
}

func MorePractice() {
	names := []string{"Diksha", "Sheetal", "Praphul", "Jignesh", "Sakshi", "Thakur"}
	for _, name := range names { // index, value
		if strings.Contains(name, "Diksha") {
			msg := fmt.Sprintf("%v is Present in %v", name, names)
			fmt.Println(msg)
		}
	}
}

func Factorial(number int) int {
	if number == 0 || number == 1 {
		return 1
	}
	return number * Factorial(number-1)
}

func VariadicInput(nums ...int) int {
	var sum int = 0
	// for _, num := range nums {
	// 	sum += num
	// }
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

func SlicesGrowth() {
	nums := make([]int, 2, 3)
	for i := range 1000 {
		nums = append(nums, i)
		fmt.Printf("After appending %d: len=%d, cap=%d\n", i, len(nums), cap(nums))
	}
}
