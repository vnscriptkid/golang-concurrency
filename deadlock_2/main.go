package main

// 1. Why does it cause deadlock
func main() {
	channel := make(chan int)

	<-channel
}
