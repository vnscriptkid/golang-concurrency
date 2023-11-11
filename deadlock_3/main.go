package main

import "fmt"

// 1. Why does it cause deadlock?
// 2. How to fix it?
func main() {
	channel := make(chan int)

	go func(channel chan int) {
		for i := 0; i < 100; i++ {
			channel <- i
		}
	}(channel)

	for v := range channel {
		fmt.Println(v)
	}
}
