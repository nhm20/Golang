package main

import (
	"fmt"
	"time"
)

func main() {
	i := 4
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("default")
	}
	//no need for break statement in switch case

	//multiple condition switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	//type switch
	whoAmI := func(i interface{}) {
		switch t:=i.(type) {
		case string:
			fmt.Println("I am a string")
		case int:
			fmt.Println("I am an int")
		default:
			fmt.Println("I am something else",t)
		}
	}
	whoAmI("hello")

}
