package main

import "fmt"

func main() {
	age := 30
	if age > 18 {
		fmt.Println("You are a minor")
	} else {
		fmt.Println("You are an adult")
	}

	// if else if else
	if age < 18 {
		fmt.Println("You are a minor")
	} else if age > 18 && age < 30 {
		fmt.Println("You are a young adult")
	} else if age > 30 && age < 50 {
		fmt.Println("You are a middle aged adult")
	} else {
		fmt.Println("You are a senior citizen")
	}

	var role = "admin"
	var hasAccess = true
	if role == "admin" && hasAccess == true {
		fmt.Println("You have access to the admin panel")
	} else if role == "user" || hasAccess == true {
		fmt.Println("You have access to the user panel")
	} else {
		fmt.Println("You do not have access to the admin panel")
	}

	//we can declare variables inside if statement
	if age := 30; age > 18 {
		fmt.Println("You are a minor")
	} else if age > 30 {
		fmt.Println("You are a middle aged adult")
	}
	//go does not have ternary operator like c++

}
