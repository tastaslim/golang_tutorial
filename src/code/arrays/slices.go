package arrays

import (
	"fmt"
	"slices"
)

func SlicesPractice() {
	/*
		// Slices are dynamics array
		var marks []int
		if marks == nil {
			fmt.Printf("Empty Slice with length %v\n", len(marks))
		}
		marks = append(marks, 1)
		marks = append(marks, 2)
		fmt.Println(marks)
	*/
	// Like C++, new keyword, in go we use make method to create these

	var marks []int = make([]int, 2) // [0,0] ==> Initial size
	fmt.Println(marks == nil)
	fmt.Println(marks, cap(marks))
	marks = append(marks, 1)
	marks = append(marks, 2)
	fmt.Println(marks, cap(marks))
	marks = append(marks, 3)
	fmt.Println(marks, cap(marks))

	nums := [4]int{1, 2, 3, 4}
	nums2 := [4]int{1, 2, 3, 4}
	fmt.Println(slices.Equal(nums[:], nums2[:]))

}

func SlicesCheck(names []string, targetName string) bool {
	// for i := range names {
	// 	if names[i] == targetName {
	// 		return true
	// 	}
	// }
	// return false

	if slices.Contains(names, targetName) {
		return true
	}

	slices.Sort(names)
	fmt.Println(names)
	return false
}
