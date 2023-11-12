package main

import (
	"fmt"
	"time"
)

func main() {
	goChannel := make(chan int)

	go func() {
		i := 0
		for {
			goChannel <- i
			fmt.Printf("Pushed value %v to channel\n", i)
			i++
			time.Sleep(time.Millisecond * 500)
		}
	}()

	timer := time.After(3 * time.Second)
	for {
		select {
		case res := <-goChannel:
			fmt.Printf("Read value from channel: %v\n", res)
		case <-timer:
			fmt.Println("timeout!!!")
			return
		}
	}
}
