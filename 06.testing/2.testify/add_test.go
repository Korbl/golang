package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	// Use the testify library to define assertions on the Add function.
	// Leverage the assert package: "github.com/stretchr/testify/assert"
	want := 7
	got := Add(3, 4)
	assert.Equal(t, want, got)
}
