package main

import "fmt"

func adding(a, b int) int {
	return a + b
}

func main() {
	// Create a function which takes two integers and returns the sum.
	// The function should be called from a main function and print the result.
	// Validate your solution by running "go test"
	fmt.Println(adding(21, 21))
}
