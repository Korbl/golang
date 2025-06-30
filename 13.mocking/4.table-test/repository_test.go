package repository

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateInfo(t *testing.T) {
	tests := []struct {
		Name               string
		SetupMockBehaviour func(dataStore *MockDataStore)
		ExpectedOutput     *Info
		ExpectErr          bool
	}{
		{
			Name: "Successfully create info",
			SetupMockBehaviour: func(dataStore *MockDataStore) {
				dataStore.EXPECT().Store(&Info{X: 42, Y: 42}).Return(nil)
			},
			ExpectErr: false,
			ExpectedOutput: &Info{
				X: 42,
				Y: 42,
			},
		},
		{
			Name: "Creating the info failed",
			SetupMockBehaviour: func(dataStore *MockDataStore) {
				dataStore.EXPECT().Store(&Info{X: 42, Y: 42}).Return(errors.New("something broke start to panic!!!"))
			},
			ExpectErr:      true,
			ExpectedOutput: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			dataStore := NewMockDataStore(t)

			test.SetupMockBehaviour(dataStore)

			infoHandler := NewInfoHandler(dataStore)
			result, err := infoHandler.CreateNewInfo(42, 42)
			if test.ExpectErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.Equal(t, test.ExpectedOutput, result)
				assert.NoError(t, err)
			}
		})
	}
}
