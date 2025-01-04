package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// Function to be run as a goroutine
func printNumbers(name string, limit int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes
	for i := 1; i <= limit; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(100 * time.Millisecond) // Simulate work with sleep
	}
}

func Goroutines() {
	fmt.Println("Goroutines function starts.")

	// Initialize a WaitGroup
	var wg sync.WaitGroup

	// Add the number of goroutines to the WaitGroup counter
	wg.Add(2)

	// Start the first goroutine
	go printNumbers("Goroutine 1", 5, &wg)

	// Start the second goroutine
	go printNumbers("Goroutine 2", 5, &wg)

	// Goroutines function also does work
	for i := 1; i <= 3; i++ {
		fmt.Printf("Goroutines function: %d\n", i)
		time.Sleep(150 * time.Millisecond)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Println("Goroutines function ends.")
}
