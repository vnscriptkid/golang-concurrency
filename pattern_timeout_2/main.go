package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Simulated network request channel
	networkResponse := make(chan string)

	// Simulating a network request in a goroutine
	go func() {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second) // Random delay between 0-2 seconds
		networkResponse <- "Response received"
	}()

	// Setting up a timeout of 1 second
	timer := time.After(1 * time.Second)

	// Select statement to wait on multiple channel operations
	select {
	case res := <-networkResponse:
		fmt.Println(res)
	case <-timer:
		fmt.Println("timeout")
	}
}
