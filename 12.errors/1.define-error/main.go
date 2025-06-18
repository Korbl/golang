package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("negative number %f used", e)
}

// Define a error type ErrNegativeSqrt float64
// Implement the Sqrt function such that it return the Sqrt for a
// given floating point number.
// In case the input number is negative, the function should
// return the ErrNegativeSqrt.

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
