package main

import "fmt"

func main() {
	goChannel := make(chan int, 5)
	goChannel <- 1
	goChannel <- 2
	fmt.Println("Length of the channel after adding 1 and 2 ", len(goChannel))
	fmt.Println("Capacity of the channel ", cap(goChannel))
	goChannel <- 3
	fmt.Println("Length of the channel after adding 3 ", len(goChannel))
}
