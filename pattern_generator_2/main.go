package main

import (
	"fmt"
	"sync"
)

// squareGenerator takes a list of integers, squares them in separate goroutines, and sends them via a channel.
func squareGenerator(nums []int) <-chan int {
	// Buffer the channel to the size of nums.
	out := make(chan int, len(nums))

	// Add the count of numbers to the WaitGroup.
	var wg sync.WaitGroup
	wg.Add(len(nums))

	for _, n := range nums {
		// Launch a goroutine for each number.
		go func(n int) {
			defer wg.Done()
			// Send the square of the number to the channel.
			out <- n * n
		}(n)
	}

	// Start a goroutine to close the channel once all the squares have been calculated.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	nums := make([]int, 10) // Generate numbers 0 through 9.
	for i := range nums {
		nums[i] = i
	}

	// Collect the results from the channel as they are generated.
	for sq := range squareGenerator(nums) {
		fmt.Println(sq)
	}
}
