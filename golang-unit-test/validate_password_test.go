package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected bool
		result   string
	}{
		{
			name:     "input valid",
			request:  "Tokenrahasia123$",
			expected: true,
			result:   "Password valid",
		},
		{
			name:     "input pendek",
			request:  "A1",
			expected: false,
			result:   "password minimal 8 karakter",
		},
		{
			name:     "input kosong",
			request:  "",
			expected: false,
			result:   "field tidak boleh kosong",
		},
		{
			name:     "input tidak valid",
			request:  "Tokenrahasia123",
			expected: false,
			result:   "password harus ada karakter unik",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ValidatePassword(test.request)
			assert.Equal(t, test.expected, result, test.result)
		})
	}
}
