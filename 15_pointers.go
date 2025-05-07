package main

import "fmt"

//by value
func changeNum(num int) {
	num = 20
	fmt.Println("Inside changeNum", num) // Output: 20, because num is passed by value
}

//by reference
func changeNumByRef(num *int) {
	*num = 20
	fmt.Println("Inside changeNumByRef", *num) // Output: 20, because num is passed by reference
}

func main() {

	num := 10
	changeNum(num)
	fmt.Println("After changeNum: in main", num) // Output: 10, because num is passed by value

	fmt.Println("Memory address of num:", &num) // Output: Memory address of num: 0xc00000a0b0
	changeNumByRef(&num)
	fmt.Println("After changeNumByRef: in main", num) // Output: 20, because num is passed by reference
}
