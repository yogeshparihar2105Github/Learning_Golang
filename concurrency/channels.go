package concurrency

import (
	"fmt"
	"time"
)

// Worker function that sends numbers to a channel
func worker(id int, ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Worker %d sending: %d\n", id, i)
		ch <- i                            // Send data to the channel
		time.Sleep(200 * time.Millisecond) // Simulate work
	}
	close(ch) // Close the channel after sending all data
}

// Function to receive data from a channel
func receiver(ch chan int) {
	for value := range ch {
		fmt.Printf("Receiver received: %d\n", value)
	}
	fmt.Println("Receiver finished.")
}

func Channels() {
	fmt.Println("Channels function starts.")

	// Create an unbuffered channel
	ch := make(chan int)

	// Start the worker as a goroutine
	go worker(1, ch)

	// Start the receiver as a goroutine
	go receiver(ch)

	// Give goroutines time to finish
	time.Sleep(2 * time.Second)

	fmt.Println("Channels function ends.")
}
