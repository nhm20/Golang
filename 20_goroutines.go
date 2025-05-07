package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("Task", id, "is running")
}

func main() {
	for i := 0; i <= 10; i++ {
		// go task(i)

		go func (i int) {
			fmt.Println("Task", i, "is running")
		}(i)

	}
	time.Sleep(2 * time.Second) // Wait for goroutines to finish
}
