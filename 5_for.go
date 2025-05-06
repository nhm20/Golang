package main

import "fmt"

//for -> only loop in go

func main() {
	//while loop
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}

	//infinite loop
	// for{
	// 	println("infinite loop")
	// }

	//classic for loop
	for i:=0;i<3;i++{
		fmt.Println("classic for loop", i)
	}

	//break and continue
	for i:=0;i<5;i++{
		if i == 2 {
			continue 
		}
		if i == 4 {
			break 
		}
		fmt.Println("break and continue", i)
	}

	//range
	for i:=range 3{
		fmt.Println("range", i) // 0 1 2
	}

}
