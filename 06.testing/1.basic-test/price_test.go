package price

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTotalCost(t *testing.T) {
	// Write one or more unit test cases for the following function, using Go's standard library
	// You can execute these tests using "go test".
	// Try passing -v and -cover arguments to go test.
	want := 1806
	out := CalculateTotalPrice(42, 42, 42)
	assert.Equal(t, want, out)
	if out != want {
		t.Errorf("got %d , want %d", out, want)
	}
}
