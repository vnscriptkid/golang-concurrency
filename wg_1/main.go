package main

import (
	"fmt"
)

func main() {
	// Why nothing is printed

	go func() {
		fmt.Println("go routine 1")
	}()

	go func() {
		fmt.Println("go routine 2")
	}()
}
