package main

import (
	"fmt"
	"sync"
)

func work(id int, w *sync.WaitGroup) {
	defer w.Done() // Decrement the counter when the goroutine completes
	fmt.Println("Task", id, "is running")
}

func main() {
	var wg sync.WaitGroup // Create a WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1)       // Increment the WaitGroup counter
		go work(i, &wg) // Pass the WaitGroup to the goroutine

	}
	wg.Wait() // Wait for all goroutines to finish
}
