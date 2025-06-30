package main

import "fmt"

func main() {
	ch := make(chan int)
	go func(a, b int, ch chan<- int) {
		c := a + b
		ch <- c
	}(1, 2, ch)
	// Retrieve the computed value c from the goroutine
	// and print it in the main function.
	// Hint: use a channel to retrieve the value.
	received := <-ch
	fmt.Printf("Received %d\n", received)
}
