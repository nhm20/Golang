package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processNum(numChannel chan int) {
	fmt.Println("Processing number...", <-numChannel) // Receive a number from the channel

}

// sending
func processNum2(numChannel chan int) {
	for num := range numChannel { // Receive numbers from the channel until it's closed
		fmt.Println("Processing number...", num) // Process the number
		time.Sleep(time.Second)
	}
}

// receiving
func sum(result chan int, num1 int, num2 int) {
	sum := num1 + num2
	result <- sum

}

// goroutine synchronizer
func tasks(done chan bool) {
	defer func() { done <- true }() // Notify that the task is done
	fmt.Println("Task is done")
}

// sending and receiving are blocking operations
func main() {

	// messageChannel := make(chan string) // Create a channel to communicate messages

	// messageChannel <- "Hello, World!" // Send a message to the channel
	// //channels are blocking

	// message := <-messageChannel // Receive the message from the channel
	// fmt.Println(message)        // Print the received message

	numChannel := make(chan int) // Create a channel to communicate numbers
	go processNum(numChannel)
	numChannel <- 42 // Send a number to the channel

	time.Sleep(2 * time.Second)

	numChannel2 := make(chan int) // Create a channel to communicate numbers
	go processNum2(numChannel2)
	for {
		numChannel2 <- rand.Intn(100)
	}

	result := make(chan int)
	go sum(result, 10, 20)
	res := <-result
	fmt.Println("Sum:", res)

	done := make(chan bool) // Create a channel to signal when the task is done - unbuffered channel
	go tasks(done)          // Start the task in a goroutine
	<-done                  // block until the task is done

	// buffered channel
	emailChan := make(chan string, 100) // Create a buffered channel with a capacity of 100 - unblocked

	emailChan <- "1@example.com"
	emailChan <- "2@example.com"
	fmt.Println((<-emailChan)) // Receive a message from the channel
	fmt.Println((<-emailChan)) // Receive a message from the channel
}
