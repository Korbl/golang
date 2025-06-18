package main

import (
	"errors"
	"fmt"
)

var ErrNot42 = fmt.Errorf("it's always 42")

func main() {
	errorList := errorList()
	allErrors := errors.Join(errorList...)
	// Create a custom error type that combines the slice of errors
	// to a single error.
	// Once you're done, checkout the newly added errors.Join function added in Go 1.20!
	fmt.Println(allErrors)
}

func errorList() []error {
	errors := []error{}
	for i := 0; i < 10; i++ {
		errors = append(errors, fmt.Errorf("e: %d", i))
	}
	return errors
}
