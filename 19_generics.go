package main

import "fmt"

func printSlice[T any](arr []T) {
	for _, item := range arr {
		fmt.Println(item)
	}
}
func printSlice2[T interface{}](arr []T) {
	for _, item := range arr {
		fmt.Println(item)
	}
}
func printSlice3[T int | string](arr []T) {
	for _, item := range arr {
		fmt.Println(item)
	}
}

func printSlice4[T comparable, V string](arr []T, str V) {
	for _, item := range arr {
		fmt.Println(item, str)
	}
}

type stack[T any] struct {
	elements []T
}

func main() {
	printSlice([]int{1, 2, 3, 4, 5})
	printSlice([]string{"a", "b", "c", "d", "e"})

	printSlice2([]float32{1.1, 2.2, 3.3, 4.4, 5.5})

	myStack := stack[int]{
		elements: []int{1, 2, 3, 4, 5},
	}
	fmt.Println(myStack)

}
