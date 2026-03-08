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
	if marks == nil {                // Now it won't be nil
		fmt.Printf("Empty Slice with length %v\n", len(marks))
	}
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
