package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	result := sum(1, 2, 3, 4, 5)
	fmt.Println(result) // Output: 15

	nums := []int{1, 2, 3, 4, 5}
	result = sum(nums...) // Unpacking the slice
	fmt.Println(result)   // Output: 15
}
