package main

import "fmt"

func main() {
	var name string = "john"
	fmt.Println(name)

	var name2 = "john" // type inference
	fmt.Println(name2)

	var isAdult bool = true
	fmt.Println(isAdult)

	var age int = 30
	fmt.Println(age)

	//shortcut syntax
	name3 := "john"
	fmt.Println(name3)

	var name4 string
	name4 = "john"
	fmt.Println(name4)

	var price float64 = 10.5
	fmt.Println(price)
}