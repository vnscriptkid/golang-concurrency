package main

import (
	"context"
	"fmt"
	"time"
)

// simulateWork simulates a long-running function that can be interrupted.
func simulateWork(ctx context.Context) error {
	// This loop represents the work being done.
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Working...")
		case <-ctx.Done():
			fmt.Println("simulateWork: Done!")
			// If the context is cancelled, return the context's error.
			return ctx.Err()
		}
	}
}

func main() {
	// Set a timeout for the simulateWork function.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // It's good practice to call cancel even if the context expires.

	// Run the function in a goroutine so that main can block on ctx.Done.
	go func() {
		err := simulateWork(ctx)
		if err != nil {
			fmt.Printf("simulateWork stopped: %v\n", err)
		}
	}()

	// Wait for the work to be done or for the timeout to expire.
	fmt.Println("Waiting for work to finish...")
	<-ctx.Done()
	fmt.Println("main: Done!")

	// Check why the context was cancelled (whether it was due to the timeout or not).
	switch ctx.Err() {
	case context.DeadlineExceeded:
		fmt.Println("The function timed out!")
	case context.Canceled:
		fmt.Println("The context was cancelled.")
	}
}
