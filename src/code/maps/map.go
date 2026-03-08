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
