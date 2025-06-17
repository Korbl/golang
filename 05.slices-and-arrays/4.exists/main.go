package main

import "slices"

func Exists(input []int, x int) bool {
	return slices.Contains(input, x)
}

func main() {
	// Define a function Exists, which takes a slice of integers and a single integer
	// and reports if the given integer exists in the slice by returning a boolean.
	// run the test to verify your answer (go test)
}
