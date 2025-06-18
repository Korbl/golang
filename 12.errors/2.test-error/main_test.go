package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	// Define multiple test cases using table tests.
	tests := []struct {
		Name   string
		X      float64
		Output float64
	}{
		{
			Name:   "Positive numbers",
			X:      4,
			Output: 2,
		},
		{
			Name:   "Negative number",
			X:      -3,
			Output: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			out, err := Sqrt(test.X)
			if err == nil {
				assert.Equal(t, test.Output, out)
			} else {
				var validationErr ErrNegativeSqrt
				assert.ErrorAs(t, err, &validationErr)
			}
		})
	}
}
