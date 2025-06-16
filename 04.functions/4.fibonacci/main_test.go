package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonnaci(t *testing.T) {
	assert.Equal(t, 1, fib(1))
	assert.Equal(t, 13, fib(7))
	assert.Equal(t, 6765, fib(20))
}
