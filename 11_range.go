package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(nums); i++ {
		println(nums[i])
	}

	sum := 0
	for idx, num := range nums {
		sum += num
		println(idx) // index
		println(num)
	}
	println(sum) // 15

	m := map[string]string{"name": "John", "age": "30"}
	for k, v := range m {
		fmt.Println(k, v) // name John age 30
	}
	for k:= range m {
		fmt.Println(k) // name age
	}

	//unicode code point rune
	//starting byte of rune
	//255-> 1 byte
	for i,c:= range "golang" {
		fmt.Println(i, c) // 0 103 1 111 2 108 3 97 4 110 5 103
		fmt.Println(string(c)) // g o l a n g
	}
}
