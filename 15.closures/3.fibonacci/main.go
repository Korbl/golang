package main

import "fmt"

// implement the fibonacci using a closure.
func fib() func() int {
	firstValue := 0
	secondValue := 1
	return func() int {
		newValue := firstValue + secondValue
		firstValue = secondValue
		secondValue = newValue
		return newValue
	}
}

func main() {
	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
