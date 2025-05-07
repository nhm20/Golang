package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
func subtract(a, b int) int { // Function with multiple parameters of the same type
	return a - b
}

func getLanguages() (string, string, bool) {
	return "Go", "Python", true
}

func processIt(fn func(a int) int) {
	fn(1)
}
func processIt2() func(a int) int {
	return func(a int) int {
		return a + 1
	}
}

func main() {
	result := add(1, 2)
	fmt.Println(result) // Output: 3
	result = subtract(5, 3)
	fmt.Println(result) // Output: 2
	lang1, lang2, lang3 := getLanguages()
	fmt.Println(lang1, lang2, lang3)
	lang1, lang2, _ = getLanguages() // Ignore the third return value
	fmt.Println(lang1, lang2)        // Output: Go Python

	fn := func(a int) int {
		return a + 1
	}
	processIt(fn)

	fmt.Println(fn(1)) // Output: 2

	fn2 := processIt2()
	fmt.Println(fn2(1)) // Output: 2
}
