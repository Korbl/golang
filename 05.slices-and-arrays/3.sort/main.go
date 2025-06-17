package main

import (
	"sort"
)

func Sort(input []string) []string {
	sort.Strings(input)
	return input
}

func main() {
	// Define a function Sort, which takes a slice of strings and orders them
	// alphabetically and returns the result. Tip: use the "sort" package

	// Bonus: try the same exercise with the "slices" package

	// run the test to verify your answer (go test)
	sortedSlice := Sort([]string{"ab", "aa", "aaa"})
	println(sortedSlice)
}
