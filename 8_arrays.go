package main

import "fmt"

// Arrays are fixed-size collections of elements of the same type.
func main() {
	var nums [4]int
	fmt.Println((len(nums)))
	nums[0] = 1
	nums[1] = 2
	fmt.Println(nums[0])
	fmt.Println(nums)

	var vals [4]bool
	vals[3] = true
	fmt.Println(vals)

	var names [4]string
	names[0] = "John"
	fmt.Println(names)

	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	//2d array
	arr2:=[2][2]int{{1, 2}, {3, 4}}
	fmt.Println(arr2)

	//-fixed size, that is predictable
	//- memory optimization, as the size is known at compile time
	// constant time access
}
