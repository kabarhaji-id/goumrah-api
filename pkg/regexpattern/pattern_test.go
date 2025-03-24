package regexpattern_test

import (
	"testing"

	"github.com/kabarhaji-id/goumrah-api/pkg/regexpattern"
)

func TestPhoneNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "10 digit phone number",
			input:    "0812345678",
			expected: true,
		},
		{
			name:     "11 digit phone number",
			input:    "08123456789",
			expected: true,
		},
		{
			name:     "12 digit phone number",
			input:    "081234567890",
			expected: true,
		},
		{
			name:     "13 digit phone number",
			input:    "0812345678901",
			expected: true,
		},
		{
			name:     "with country code phone number",
			input:    "+6281234567890",
			expected: true,
		},
		{
			name:     "11 digit with country code phone number",
			input:    "+62812345678",
			expected: true,
		},
		{
			name:     "12 digit with country code phone number",
			input:    "+628123456789",
			expected: true,
		},
		{
			name:     "13 digit with country code phone number",
			input:    "+62812345678901",
			expected: true,
		},
		{
			name:     "without prefix phone number",
			input:    "81234567890",
			expected: false,
		},
		{
			name:     "less than 10 digit phone number",
			input:    "081234567",
			expected: false,
		},
		{
			name:     "more than 13 digit phone number",
			input:    "08123456789012",
			expected: false,
		},
		{
			name:     "with hyphen phone number",
			input:    "0812-3456-7890",
			expected: false,
		},
		{
			name:     "with space phone number",
			input:    "0812 3456 7890",
			expected: false,
		},
		{
			name:     "with non-8 after prefix phone number",
			input:    "09123456789",
			expected: false,
		},
		{
			name:     "with alpha phone number",
			input:    "0812abc6789",
			expected: false,
		},
		{
			name:     "empty",
			input:    "",
			expected: false,
		},
		{
			name:     "only country code",
			input:    "+62",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := regexpattern.PhoneNumber().MatchString(test.input)
			if result != test.expected {
				t.Fatalf("expected %v, got %v", test.expected, result)
			}
		})
	}
}
