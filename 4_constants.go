package main

import "fmt"

const age = 30

var name = "John " // this is a global variable
func main() {
	const name = "John Doe"

	fmt.Println(name) 
	
	const (
		port=8080
		host="localhost"
	)
	fmt.Println(port, host) // 8080 localhost
}