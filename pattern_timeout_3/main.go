package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	clientRequests := make(chan string)

	// Simulate client requests
	go func() {
		for {
			time.Sleep(time.Duration(2+rand.Intn(5)) * time.Second)
			clientRequests <- "Client Request"
		}
	}()

	for {
		select {
		case req := <-clientRequests:
			fmt.Printf("Received: %s\n", req)
			// Process the request
		case <-time.After(5 * time.Second):
			fmt.Println("No requests for 5 seconds, checking system health...")
			// Perform timeout action, like checking system health
		}
	}
}
