package arrays

import (
	"fmt"
	"slices"
	"strings"

	"github.com/tastaslim/golang_tutorial/src/pkg/logger"
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

func BubbleSort(arr []int) []int {
	sizeOfArray := len(arr)
	for i := range sizeOfArray {
		for j := range i {
			tempElement := arr[i]
			arr[i] = arr[j]
			arr[j] = tempElement
		}
	}
	return arr
}

func RevereseArray(arr []int) []int {
	endIndex := len(arr) - 1
	for startIndex := range endIndex {
		tempElement := arr[endIndex]
		arr[endIndex] = arr[startIndex]
		arr[startIndex] = tempElement
		startIndex += 1
		endIndex -= 1
	}

	return arr
}

type TupleSum struct {
	firstNumber  int
	secondNumber int
}

func PairSum(arr []int, target int) []TupleSum {
	endIndex := len(arr) - 1
	startIndex := 0
	var finalTuple []TupleSum
	for startIndex < endIndex {
		currentSum := arr[startIndex] + arr[endIndex]
		if currentSum == target {
			finalTuple = append(finalTuple, TupleSum{firstNumber: arr[startIndex], secondNumber: arr[endIndex]})
			startIndex += 1
			endIndex -= 1
		} else if currentSum < target {
			startIndex += 1
		} else {
			endIndex -= 1
		}
	}
	return finalTuple
}

func UnionArray(arr1 []int, arr2 []int) []int {
	var finalArray []int
	start1, end1, start2, end2 := 0, len(arr1), 0, len(arr2)
	for start1 < end1 && start2 < end2 {
		if arr1[start1] < arr2[start2] {
			finalArray = append(finalArray, arr1[start1])
			start1 += 1
		} else if arr1[start1] > arr2[start2] {
			finalArray = append(finalArray, arr2[start2])
			start2 += 1
		} else {
			for arr1[start1] == arr2[start2] {
				finalArray = append(finalArray, arr2[start2])
				start1 += 1
				start2 += 1
			}
		}
	}

	for start1 < end1 {
		finalArray = append(finalArray, arr1[start1])
		start1 += 1
	}

	for start2 < end2 {
		finalArray = append(finalArray, arr2[start2])
		start2 += 1
	}

	return finalArray
}

func ElementWithMaxFrequency(arr []int) int {
	hashMap := make(map[int]int)
	sizeOfArray := len(arr)
	maxFrequencyElement := arr[0]
	maxFrequency := 0
	for i := range sizeOfArray {
		hashMap[arr[i]]++
		if hashMap[arr[i]] > maxFrequency {
			maxFrequencyElement = arr[i]
			maxFrequency = hashMap[arr[i]]
		}
	}
	return maxFrequencyElement
}

/*
Given an array of positive integers arr[], return the second largest element from the array. If the second largest element doesn't exist then return -1.

Note: The second largest element should not be equal to the largest element.

Examples:

Input: arr[] = [12, 35, 1, 10, 34, 1]
Output: 34
Explanation: The largest element of the array is 35 and the second largest element is 34.
Input: arr[] = [10, 5, 10]
Output: 5
Explanation: The largest element of the array is 10 and the second largest element is 5.
Input: arr[] = [10, 10, 10]
Output: -1
Explanation: The largest element of the array is 10 and the second largest element does not exist.
Constraints:
2 ≤ arr.size() ≤ 105
1 ≤ arr[i] ≤ 105
*/

func SecondLargestElement(arr []int) int {
	largest, secondLargest := arr[0], -1
	sizeOfArray := len(arr)
	for i := range sizeOfArray {
		if arr[i] > largest {
			secondLargest = largest
			largest = arr[i]
		} else if arr[i] > secondLargest && arr[i] != largest {
			secondLargest = arr[i]
		}
	}
	return secondLargest
}

func MaximumElement(arr []int) int {
	maximumElement := 0
	sizeOfArray := len(arr)
	for i := range sizeOfArray {
		maximumElement = max(maximumElement, arr[i])
	}
	return maximumElement
}

func InitializeArray(size int) []int {
	employees := [10]string{}
	for range employees {
		TakeInput()
		employees = append(employees, element)
	}
}
