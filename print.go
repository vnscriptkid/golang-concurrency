package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func lowerCase(goChannel chan<- string) {
	defer wg.Done()
	for chr := 'a'; chr <= 'z'; chr++ {
		goChannel <- fmt.Sprintf("%c", chr)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
	}
}

func upperCase(goChannel chan<- string) {
	defer wg.Done()
	for chr := 'A'; chr <= 'Z'; chr++ {
		goChannel <- fmt.Sprintf("%c", chr)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Microsecond)
	}
}

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(4)
	wg.Add(2)
	goChannel := make(chan string)
	go lowerCase(goChannel)
	go upperCase(goChannel)

	go func() {
		for val := range goChannel {
			fmt.Println(val)
		}
	}()

	wg.Wait()
}
