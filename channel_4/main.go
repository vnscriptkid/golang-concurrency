package main

import "fmt"

func next(n int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		i, j := 0, 1
		for k := 0; k < n; k++ {
			out <- i
			i, j = i+j, i
		}
	}()
	return out
}

// pass unidirectional channel and number of Fibonacci values required as input
// pass value received from the next function to the channel
func useFibo(sendChan chan<- int, n int) {
	for x := range next(n) {
		sendChan <- x
	}
}

func main() {
	// number of Fibonacci required
	n := 10
	// create channel
	channel := make(chan int)
	// initialize the function
	go useFibo(channel, n)
	// print the values
	for i := 0; i < 10; i++ {
		fmt.Println("Received: ", <-channel)
	}
}
