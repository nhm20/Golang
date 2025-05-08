package main

import (
	"fmt"
	"time"
)

func emailSender2(emailChan chan string, done chan bool) {
	defer func() { done <- true }() // Notify that the task is done
	for email := range emailChan {  // Receive emails from the channel until it's closed
		fmt.Println("Sending email to:", email) // Process the email
		time.Sleep(time.Second)                 // Simulate email sending delay
	}
}

func main() {
	// buffered channel
	emailChan := make(chan string, 5) // Create a buffered channel with a capacity of 100 - unblocked

	done := make(chan bool) // Create a channel to signal when the task is done - unbuffered channel

	// emailChan <- "1@example.com"
	// emailChan <- "2@example.com"
	// fmt.Println((<-emailChan))
	// fmt.Println((<-emailChan))

	go emailSender2(emailChan, done) // Start the email sender in a goroutine

	for i := 0; i < 5; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i) //unblocking
	}
	fmt.Println("Emails sent")
	close(emailChan) // Close the channel to signal that no more emails will be sent
	<-done
}
