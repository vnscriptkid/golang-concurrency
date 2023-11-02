package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

var m map[string]int = map[string]int{
	"key": 100,
}

func read() {
	for {
		mu.Lock()
		fmt.Println(m["key"])
		mu.Unlock()
	}
}

func write() {
	for {
		mu.Lock()
		m["key"] = m["key"] + 1
		mu.Unlock()
	}
}

func main() {
	forever := make(chan bool)

	go read()
	go write()

	<-forever
}
