package main

import (
	"fmt"
	"sync"
)

var x int = 0

var wg sync.WaitGroup

// add 100
func increment() {
	defer wg.Done()
	for i := 0; i < 10000000; i++ {
		x = x + 1
	}
}

func main() {
	wg.Add(2)
	// Expect +10000000 +10000000 => 20000000
	go increment()
	go increment()
	wg.Wait()

	fmt.Println("x = ", x)
}
