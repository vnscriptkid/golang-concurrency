package main

import (
	"fmt"
	"sort"
	"sync"
)

var wg sync.WaitGroup

func main() {
	output := make([]int, 0)

	for val := range producer() {
		output = append(output, val)
	}

	// sort output in asc order using go standard library
	sort.Ints(output)

	fmt.Println(output)
}

func producer() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 100; i += 5 {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				out <- i
			}(i)
		}
		wg.Wait()
		close(out)
	}()
	return out
}
