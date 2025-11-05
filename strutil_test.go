package cpfcnpj

import (
	"testing"
)

func TestClean(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "CPF with dots and dash",
			input:    "111.444.777-35",
			expected: "11144477735",
		},
		{
			name:     "CNPJ with dots, slash and dash",
			input:    "11.222.333/0001-81",
			expected: "11222333000181",
		},
		{
			name:     "Already clean CPF",
			input:    "11144477735",
			expected: "11144477735",
		},
		{
			name:     "Already clean CNPJ",
			input:    "11222333000181",
			expected: "11222333000181",
		},
		{
			name:     "With spaces",
			input:    "111 444 777 35",
			expected: "11144477735",
		},
		{
			name:     "Mixed special characters",
			input:    "111.444-777/35",
			expected: "11144477735",
		},
		{
			name:     "With parentheses",
			input:    "(111)444-777-35",
			expected: "11144477735",
		},
		{
			name:     "With brackets",
			input:    "[111]444[777]35",
			expected: "11144477735",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Only special characters",
			input:    ".-/() []",
			expected: "",
		},
		{
			name:     "Only digits",
			input:    "1234567890",
			expected: "1234567890",
		},
		{
			name:     "With letters",
			input:    "111abc444def777ghi35",
			expected: "11144477735",
		},
		{
			name:     "Unicode characters",
			input:    "111★444★777★35",
			expected: "11144477735",
		},
		{
			name:     "Tabs and newlines",
			input:    "111\t444\n777\r35",
			expected: "11144477735",
		},
		{
			name:     "Multiple consecutive special characters",
			input:    "111...444---777///35",
			expected: "11144477735",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Clean(tt.input)
			if result != tt.expected {
				t.Errorf("Clean(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCleanRegex(t *testing.T) {
	t.Run("Regex removes non-digits", func(t *testing.T) {
		testStr := "abc123def456"
		result := Clean(testStr)
		expected := "123456"

		if result != expected {
			t.Errorf("Clean(%q) = %q, want %q", testStr, result, expected)
		}
	})

	t.Run("Regex preserves all digits", func(t *testing.T) {
		testStr := "0123456789"
		result := Clean(testStr)

		if result != testStr {
			t.Errorf("Clean(%q) = %q, want %q", testStr, result, testStr)
		}
	})
}
