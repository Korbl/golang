package main

import (
	"slices"
	"strings"
)

func Sort(input []string) []string {
	return slices.SortFunc(names, func(a, b string) int { return strings.Compare(strings.ToLower(a), strings.ToLower(b)) })
}

func main() {
	// Define a function Sort, which takes a slice of strings and orders them
	// alphabetically and returns the result. Tip: use the "sort" package

	// Bonus: try the same exercise with the "slices" package

	// run the test to verify your answer (go test)
}
