package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("go routine 1")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 1)

		fmt.Println("go routine 2")
	}()

	fmt.Println("Before wait")
	wg.Wait()
	fmt.Println("After wait")
}
