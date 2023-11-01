package main

// 1. How `select` is executed?
// - what case is chosen?
// - when default case is executed?
// - what happens if default case is added in this case?
// 2. Why does it cause deadlock
func main() {
	channel := make(chan int)

	select {
	case <-channel:
	}
}
