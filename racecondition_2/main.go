package main

import (
	"fmt"
)

var m map[string]int = map[string]int{
	"key": 100,
}

func read() {
	for {
		fmt.Println(m["key"])
	}
}

func write() {
	for {
		m["key"] = m["key"] + 1
	}
}

func main() {
	forever := make(chan bool)

	go read()
	go write()

	<-forever
}
