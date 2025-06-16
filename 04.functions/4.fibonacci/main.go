package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	// Implement the fibonacci sequence in a function, taking an integer and returning one.
	// To verify your solution run the tests.
	for i := 0; i < 20; i++ {
		fmt.Println("Fibonacci output for i = ", i, " is ", fib(i))
	}
}
