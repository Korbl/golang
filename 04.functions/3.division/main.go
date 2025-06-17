package main

import (
	"fmt"
)

func divide(x, y int) int {
	if y == 0 {
		return 0
	}
	return x / y
}

func main() {
	// Create a function which takes two 64 bit floating points as arguments.
	// The function should divide the first number by the second number.
	// Run the tests to verify if your solution is correct.
	// Note: The test assumes that you'll handle division by zero.
	// In this case return math.NaN().
	fmt.Println(divide(42, 6))
}
