package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Implement the mock behaviour
// First run `make mocks` and see what has been generated.
// Afterwards define the mock behaviour
// Hint: The testify mocks have a "On" function and a "Return function"
// Details on how the mock looks like and how they can be used,
// can be found in the testify readme: https://github.com/stretchr/testify#mock-package
func TestCreateInfo(t *testing.T) {
	dataStore := NewMockDataStore(t)
	// Define the mock behaviour
	dataStore.Store(&Info{42, 42, 42})

	infoHandler := NewInfoHandler(dataStore)
	result, err := infoHandler.CreateNewInfo(42, 42)
	assert.NoError(t, err)
	assert.Equal(t, &Info{
		X: 42,
		Y: 42,
	}, result)
}
