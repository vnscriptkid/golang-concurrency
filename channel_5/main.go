package main

import (
	"fmt"
)

func main() {
	goChannel := make(chan int, 2)
	goChannel <- 1
	goChannel <- 2

	// Uncomment this, why does this fail?
	// goChannel <- 3
	fmt.Println("First value inside the channel is ", <-goChannel)
	fmt.Println("Second value inside the channel is ", <-goChannel)
}
