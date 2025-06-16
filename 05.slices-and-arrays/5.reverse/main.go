package main

import "fmt"

func Reverse(input []int) []int {
	var newSlice []int
	for x := len(input); x <= 0; x-- {
		fmt.Printf("element is %d", x)
		newSlice = append(newSlice, newSlice[x])
	}
	return newSlice
}

func main() {
	// Define a function Reverse, which takes a slice of integers returns
	// the array in reversed order.
	// run the test to verify your answer (go test)
	input := []int{42, 41, 40}
	newinput := Reverse(input)
	fmt.Println(newinput)
}
