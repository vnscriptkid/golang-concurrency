package main

import "fmt"

func main() {
	channel := make(chan int)

	go func() {
		// Send value to channel
		channel <- 10
	}()

	// Read from channel
	fmt.Println("Received: ", <-channel)
}
