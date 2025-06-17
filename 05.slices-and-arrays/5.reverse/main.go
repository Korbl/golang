package main

import "fmt"

func Reverse(input []int) []int {
	var newSlice []int
	for x := len(input) - 1; x >= 0; x-- {
		newSlice = append(newSlice, input[x])
	}
	return newSlice
}

func main() {
	// Define a function Reverse, which takes a slice of integers returns
	// the array in reversed order.
	// run the test to verify your answer (go test)
	input := []int{42, 41, 40}
	newInput := Reverse(input)
	fmt.Println(newInput)
}
