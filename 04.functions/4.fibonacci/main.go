package main

import "fmt"

func fibonacci(i int) int {
	return (i - 1) + (i - 2)
}

func main() {
	// Implement the fibonacci sequence in a function, taking an integer and returning one.
	// To verify your solution run the tests.
	for i := 0; i < 20; i++ {
		fmt.Println("Fibonacci output for i = ", i, " is ", fibonacci(i))
	}
}
