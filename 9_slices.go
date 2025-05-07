package main

import (
	"fmt"
	"slices"
)

//slices are dynamic-size collections of elements of the same type.
//most used construct in go
// + useful methods
func main() {
	// uninitialized slice is nil
	var nums []int
	fmt.Println(nums) // nil
	fmt.Println(nums == nil) // true
	fmt.Println(len(nums)) // 0

	var nums2=make([]int, 4) // initialized slice with length 4
	fmt.Println(nums2) // [0 0 0 0]
	fmt.Println(nums2 == nil) // false

	//capacity -> max nums of elements can fit
	fmt.Println(cap(nums2)) // 4

	var nums3=make([]int, 4, 5) // initialized slice with length 4 and capacity 5
	fmt.Println(cap(nums3)) // 5

	nums3 = append(nums3, 1)
	nums3[0] = 1
	nums3 = append(nums3, 2)
	nums3 = append(nums3, 3)
	nums3 = append(nums3, 4)
	nums3 = append(nums3, 5)

	var nums4=make([]int, 0, 5) // initialized slice with length 0 and capacity 5
	fmt.Println(nums4) // []
	nums4 = append(nums4, 1)
	nums4 = append(nums4, 2)
	fmt.Println(nums4) // [1 2]

	nums5:=[]int{}
	nums5 = append(nums5, 1)
	nums5 = append(nums5, 2)
	fmt.Println(nums5) // [1 2]


	//copy function
	var nums6=make([]int,0,5)
	nums6=append(nums6, 1)
	var nums7=make([]int,len(nums6))
	
	copy(nums7, nums6) // copy nums6 to nums7
	fmt.Println(nums7) // [1 0 0 0 0]

	//slice operator
	nums8:=[]int{1,2,3,4,5}
	fmt.Println(nums8[0:2])
	fmt.Println(nums8[:1]) 
	fmt.Println(nums8[1:])

	//slices
	var nums9=[]int {1,2,3}
	var nums10=[]int{1,2,3}
	fmt.Println((slices.Equal(nums9, nums10))) // true

	var nums11=[][]int{{1,2,3},{4,5,6}}
	fmt.Println(nums11) // [[1 2 3] [4 5 6]]
	
}
