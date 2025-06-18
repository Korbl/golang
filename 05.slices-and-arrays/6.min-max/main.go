package main

func MinMax(input []int) (int, int) {
	if len(input) == 0 {
		return 0, 0
	}
	min := input[0]
	max := input[len(input)-1]
	for x := range input {
		if min > x {
			min = x
		}
		if max < x {
			max = x
		}

	}
	return min, max
}

func main() {
	// Define a function MinMax, which takes a slice of integers returns
	// the min and max values.
	// If the array is empty, the function should return min = 0, max = 0.

	// run the test to verify your answer (go test)
}
