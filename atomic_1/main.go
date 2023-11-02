package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x int64 = 0

var wg sync.WaitGroup

// add 100
func increment() {
	defer wg.Done()
	for i := 0; i < 10000000; i++ {
		// x = x + 1
		atomic.AddInt64(&x, 1)
	}
}

func main() {
	wg.Add(2)
	// Expect +10000000 +10000000 => 20000000
	go increment()
	go increment()
	wg.Wait()

	fmt.Println("x = ", atomic.LoadInt64(&x))
}
