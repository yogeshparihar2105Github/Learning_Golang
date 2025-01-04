package concurrency

import (
	"context"
	"fmt"
	"time"
)

func workerWithTimeout(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: Stopping, reason: %s\n", id, ctx.Err())
			return
		default:
			fmt.Printf("Worker %d: Working...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func Context() {
	fmt.Println("Context function starts.")

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // Ensure cancel is called to release resources

	// Start a worker goroutine
	go workerWithTimeout(ctx, 1)

	// Wait for the context timeout to trigger
	time.Sleep(4 * time.Second)

	fmt.Println("Context function ends.")
}
