package maps

import (
	"fmt"
)

/*
Maps in go are equivalent to dict in python and hashmap in c++
*/

func MapsPractice() {
	numFrequency := map[int]int{1: 10, 2: 5, 5: 4}
	studentNameAgeMapper := map[string]int{} // make(map[string]int)
	studentNameAgeMapper["Taslim"] = 26

	fmt.Println(studentNameAgeMapper["Taslim"])
	fmt.Println(numFrequency)
}

func MaxFrequencyElement(arr []int) (int, int) {
	frequencyCounter := make(map[int]int)
	maxCounter := 0
	maxFreqElement := arr[0]
	for i := range arr {
		frequencyCounter[arr[i]]++
		if frequencyCounter[arr[i]] > maxCounter {
			maxCounter = frequencyCounter[arr[i]]
			maxFreqElement = arr[i]
		}
	}
	return maxCounter, maxFreqElement
}

func CountFrequency(arr []int) map[int]int {
	frequencyCounter := make(map[int]int)
	lengthOfArray := len(arr)
	for i := range lengthOfArray {
		frequencyCounter[arr[i]]++
	}
	return frequencyCounter
}

func UnionArray(arr1 []int, arr2 []int) []int {
	hashMap := make(map[int]int)
	sizeOfArray1, sizeOfArray2 := len(arr1), len(arr2)
	for i := range sizeOfArray1 {
		hashMap[arr1[i]]++
	}
	for i := range sizeOfArray2 {
		hashMap[arr2[i]]++
	}
	finalArray := []int{}
	for key := range hashMap {
		finalArray = append(finalArray, key)
	}
	return finalArray
}

func CheckEqual(arr1 []int, arr2 []int) bool {
	hashMap := make(map[int]int)
	sizeOfArray1, sizeOfArray2 := len(arr1), len(arr2)
	for i := range sizeOfArray1 {
		hashMap[arr1[i]]++
	}
	for i := range sizeOfArray2 {
		count := hashMap[arr2[i]]
		if count > 1 {
			hashMap[arr2[i]]--
		} else {
			delete(hashMap, arr2[i])
		}
	}

	if len(hashMap) == 0 {
		return true
	}
	return false
}

func IsSubset(arr1 []int, arr2 []int) bool {
	hashMap := make(map[int]int)
	sizeOfArray1, sizeOfArray2 := len(arr1), len(arr2)
	for i := range sizeOfArray1 {
		hashMap[arr1[i]]++
	}
	for i := range sizeOfArray2 {
		hashMap[arr2[i]]--
	}

	for _, value := range hashMap {
		if value < 0 {
			return false
		}
	}

	return true
}
