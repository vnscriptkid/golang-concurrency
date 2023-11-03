package main

import (
	"fmt"
	"sync"
	"time"
)

func heavyCal(a, b int, wg *sync.WaitGroup) {
	fmt.Printf("Start: %v + %v = ???\n", a, b)
	if wg != nil {
		defer wg.Done()
	}
	r := a + b
	// Simulate a heavy calculation
	time.Sleep(time.Second * 1)

	fmt.Printf("Result: %v + %v = %v\n", a, b, r)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go heavyCal(1, 1, &wg)
	go heavyCal(2, 2, &wg)

	// Anonymous goroutine
	go func(a, b int, wg *sync.WaitGroup) {
		fmt.Printf("Start: %v + %v = ???\n", a, b)
		if wg != nil {
			defer wg.Done()
		}
		r := a + b
		// Simulate a heavy calculation
		time.Sleep(time.Second * 1)

		fmt.Printf("Result: %v + %v = %v\n", a, b, r)
	}(3, 3, &wg)

	wg.Wait()
	fmt.Println("Done")
}
