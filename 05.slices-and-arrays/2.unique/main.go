package main

import (
	"slices"
)

func Unique(input []int) []int {
	var newSlice []int
	newSlice = []int{}
	for _, v := range input {
		if !slices.Contains(newSlice, v) {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func main() {
	// Define a function Unique that takes a slice of integers and returns a new slice
	// containing only unique integers. The resulting slice should remain in the same order.
	// run the test to verify your answer (go test)
}
