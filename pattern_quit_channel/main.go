package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup
	jobs := make(chan int, 10)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				done <- true
				// doing the stuff after we receive the signal
				fmt.Println("Received all jobs")
				return
			}
		}
	}()

	wg1.Add(1)
	wg2.Add(1)
	go func() {
		defer wg1.Done()
		for j := 1; j <= 100; j++ {
			jobs <- j
			fmt.Println("sending job", j)
		}
	}()

	go func() {
		defer wg2.Done()
		wg1.Wait()
		close(jobs)
		fmt.Println("Sent all jobs")
		// passing the signal to return from the goroutine
		<-done
	}()

	wg2.Wait()
}
