package main

import (
	"fmt"
	"maps"
)

func main() {
	//creating a map

	mp := make(map[string]string) // map with string keys and string values

	//setting an element
	mp["name"] = "John"
	mp["age"] = "30"

	//getting an element
	fmt.Println(mp["name"]) // John

	//IMP: if key not found, returns zero value of the value type
	fmt.Println(mp["address"]) // ""

	m:=make(map[string]int)
	m["age"] = 30
	fmt.Println(m["age"]) // 30
	fmt.Println(len(m)) // 1

	delete(m, "age") // delete key "age"
	fmt.Println(m) // map[]

	//clearing a map
	clear(m) // clear map
	fmt.Println(m) // map[]

	mp2:= map[string]string{"name": "John", "age": "30"}
	fmt.Println(mp2) // map[age:30 name:John]

	v, ok := mp2["name"] // check if key exists
	fmt.Println(v, ok) // John true
	if ok {
		fmt.Println(v) // John
	} else {
		fmt.Println("key not found")
	}

	//maps
	m1:=map[string]int{"price": 100, "quantity": 2}
	m2:=map[string]int{"price": 200, "quantity": 3}
	fmt.Println(maps.Equal(m1, m2)) 
}
