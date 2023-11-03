package main

import "fmt"

func main() {
	channel := make(chan int)

	// Send value to channel
	channel <- 10

	// Read from channel
	fmt.Println(<-channel)

	// fatal error: all goroutines are asleep - deadlock!
	// WHY?
}
