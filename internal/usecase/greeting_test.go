package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/PlatformStackPulse/go-template/internal/logger"
	"github.com/PlatformStackPulse/go-template/internal/usecase"
)

func TestGreetingUseCaseExecute(t *testing.T) {
	log := logger.NewLogger(false)

	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name:     "greeting with name",
			input:    "Charlie",
			expected: "Hello, Charlie!",
			wantErr:  false,
		},
		{
			name:     "greeting without name",
			input:    "",
			expected: "Hello, World!",
			wantErr:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			uc := usecase.NewGreetingUseCase(log)
			result, err := uc.Execute(tc.input)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}
