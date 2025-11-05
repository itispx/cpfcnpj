package cpfcnpj

import (
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedValid bool
		expectedType  string
	}{
		{
			name:          "Valid CPF without formatting",
			input:         "11144477735",
			expectedValid: true,
			expectedType:  "cpf",
		},
		{
			name:          "Valid CPF with formatting",
			input:         "111.444.777-35",
			expectedValid: true,
			expectedType:  "cpf",
		},
		{
			name:          "Valid CPF with mixed formatting",
			input:         "111.444.77735",
			expectedValid: true,
			expectedType:  "cpf",
		},
		{
			name:          "Valid CNPJ without formatting",
			input:         "11222333000181",
			expectedValid: true,
			expectedType:  "cnpj",
		},
		{
			name:          "Valid CNPJ with formatting",
			input:         "11.222.333/0001-81",
			expectedValid: true,
			expectedType:  "cnpj",
		},
		{
			name:          "Invalid CPF - all zeros",
			input:         "00000000000",
			expectedValid: false,
			expectedType:  "cpf",
		},
		{
			name:          "Invalid CPF - wrong check digits",
			input:         "11144477736",
			expectedValid: false,
			expectedType:  "cpf",
		},
		{
			name:          "Invalid CNPJ - wrong check digits",
			input:         "11222333000182",
			expectedValid: false,
			expectedType:  "cnpj",
		},
		{
			name:          "Invalid length",
			input:         "123456789",
			expectedValid: false,
			expectedType:  "",
		},
		{
			name:          "Empty string",
			input:         "",
			expectedValid: false,
			expectedType:  "",
		},
		{
			name:          "Only special characters",
			input:         "...-/",
			expectedValid: false,
			expectedType:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, docType := Validate(&ValidateInput{CPFCNPJ: tt.input})

			if valid != tt.expectedValid {
				t.Errorf("Validate() valid = %v, want %v", valid, tt.expectedValid)
			}

			if docType != tt.expectedType {
				t.Errorf("Validate() type = %v, want %v", docType, tt.expectedType)
			}
		})
	}
}

func TestMask(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		expectError bool
	}{
		{
			name:        "Mask CPF",
			input:       "11144477735",
			expected:    "111.444.777-35",
			expectError: false,
		},
		{
			name:        "Mask CPF with existing formatting",
			input:       "111.444.777-35",
			expected:    "111.444.777-35",
			expectError: false,
		},
		{
			name:        "Mask CNPJ",
			input:       "11222333000181",
			expected:    "11.222.333/0001-81",
			expectError: false,
		},
		{
			name:        "Mask CNPJ with existing formatting",
			input:       "11.222.333/0001-81",
			expected:    "11.222.333/0001-81",
			expectError: false,
		},
		{
			name:        "Invalid length - too short",
			input:       "123456789",
			expected:    "123456789",
			expectError: true,
		},
		{
			name:        "Invalid length - too long",
			input:       "123456789012345",
			expected:    "123456789012345",
			expectError: true,
		},
		{
			name:        "Empty string",
			input:       "",
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Mask(&MaskInput{CPFCNPJ: tt.input})

			if tt.expectError && err == nil {
				t.Error("Mask() expected error, got nil")
			}

			if !tt.expectError && err != nil {
				t.Errorf("Mask() unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Mask() = %v, want %v", result, tt.expected)
			}
		})
	}
}
