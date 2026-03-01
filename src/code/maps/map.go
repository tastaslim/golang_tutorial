package maps

import (
	"fmt"
)

/*
Maps in go are equivalent to dict in python and hashmap in c++
*/

func MapsPractice() {
	students := map[string]int{"Taslim": 1, "Diksha": 2, "Asad": 3}
	fmt.Println(students)
}
