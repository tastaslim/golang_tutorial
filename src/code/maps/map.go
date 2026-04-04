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
