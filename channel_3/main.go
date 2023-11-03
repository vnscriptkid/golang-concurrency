package main

import "fmt"

func fetchValue(c chan int) {
	// Send value to channel
	c <- 10

	close(c) // Implicitly send value over channel (zero-value, false)
}

func main() {
	channel := make(chan int)

	go fetchValue(channel)

	// Read from channel
	v, ok := <-channel
	fmt.Println("Received 1: ", v, ok)
	v, ok = <-channel
	fmt.Println("Received 2: ", v, ok)
}
