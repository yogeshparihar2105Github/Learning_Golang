package concurrency

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int            // Shared resource
	mutex   sync.Mutex     // Mutex for synchronizing access to the shared resource
	wg      sync.WaitGroup // WaitGroup to wait for goroutines to finish
)

// Function to increment the counter safely
func increment(id int) {
	defer wg.Done() // Notify WaitGroup that this goroutine is done
	for i := 0; i < 5; i++ {
		mutex.Lock() // Lock the shared resource
		counter++
		fmt.Printf("Goroutine %d incremented counter to: %d\n", id, counter)
		mutex.Unlock()                     // Unlock the shared resource
		time.Sleep(100 * time.Millisecond) // Simulate work
	}
}

// Singleton initialization using sync.Once
var once sync.Once
var singletonInstance *Singleton

type Singleton struct {
	Value string
}

// Function to initialize the singleton
func getInstance() *Singleton {
	once.Do(func() {
		fmt.Println("Initializing singleton instance...")
		singletonInstance = &Singleton{Value: "Singleton Value"}
	})
	return singletonInstance
}

func Sync() {
	fmt.Println("Sync function starts.")

	// Synchronize with WaitGroup
	wg.Add(3) // Add three goroutines to the WaitGroup

	// Start multiple goroutines
	for i := 1; i <= 3; i++ {
		go increment(i)
	}

	wg.Wait() // Wait for all goroutines to complete
	fmt.Println("Final Counter Value:", counter)

	// Demonstrating sync.Once for singleton pattern
	fmt.Println("\nDemonstrating sync.Once:")
	instance1 := getInstance()
	instance2 := getInstance()
	fmt.Printf("Instance1: %v, Instance2: %v\n", instance1, instance2)

	fmt.Println("Sync function ends.")
}
