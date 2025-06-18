package main

import (
	"errors"
	"fmt"
)

var ErrNot42 = fmt.Errorf("it's always 42")

func main() {
	err := ValidateNumber(44)
	for err != nil {
		fmt.Println(err)
		err = errors.Unwrap(err)
	}
}

func ValidateNumber(x int) error {
	if x != 42 {
		return fmt.Errorf("wrap tastic: %w", fmt.Errorf("start to panic: %w", ErrNot42))
	}
	return nil
}
