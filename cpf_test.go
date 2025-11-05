package cpfcnpj

import (
	"testing"
)

func TestValidateCPF(t *testing.T) {
	tests := []struct {
		name     string
		cpf      string
		expected bool
	}{
		{
			name:     "Valid CPF 1",
			cpf:      "11144477735",
			expected: true,
		},
		{
			name:     "Valid CPF 2",
			cpf:      "52998224725",
			expected: true,
		},
		{
			name:     "Valid CPF with formatting",
			cpf:      "111.444.777-35",
			expected: true,
		},
		{
			name:     "All zeros",
			cpf:      "00000000000",
			expected: false,
		},
		{
			name:     "All ones",
			cpf:      "11111111111",
			expected: false,
		},
		{
			name:     "All twos",
			cpf:      "22222222222",
			expected: false,
		},
		{
			name:     "All threes",
			cpf:      "33333333333",
			expected: false,
		},
		{
			name:     "All fours",
			cpf:      "44444444444",
			expected: false,
		},
		{
			name:     "All fives",
			cpf:      "55555555555",
			expected: false,
		},
		{
			name:     "All sixes",
			cpf:      "66666666666",
			expected: false,
		},
		{
			name:     "All sevens",
			cpf:      "77777777777",
			expected: false,
		},
		{
			name:     "All eights",
			cpf:      "88888888888",
			expected: false,
		},
		{
			name:     "All nines",
			cpf:      "99999999999",
			expected: false,
		},
		{
			name:     "Too short",
			cpf:      "1114447773",
			expected: false,
		},
		{
			name:     "Too long",
			cpf:      "111444777355",
			expected: false,
		},
		{
			name:     "Empty string",
			cpf:      "",
			expected: false,
		},
		{
			name:     "Wrong first check digit",
			cpf:      "11144477745",
			expected: false,
		},
		{
			name:     "Wrong second check digit",
			cpf:      "11144477736",
			expected: false,
		},
		{
			name:     "Both check digits wrong",
			cpf:      "11144477700",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateCPF(&ValidateCPFInput{CPF: tt.cpf})
			if result != tt.expected {
				t.Errorf("ValidateCPF(%s) = %v, want %v", tt.cpf, result, tt.expected)
			}
		})
	}
}

func TestMaskCPF(t *testing.T) {
	tests := []struct {
		name        string
		cpf         string
		expected    string
		expectError bool
	}{
		{
			name:        "Valid CPF",
			cpf:         "11144477735",
			expected:    "111.444.777-35",
			expectError: false,
		},
		{
			name:        "CPF already formatted",
			cpf:         "111.444.777-35",
			expected:    "111.444.777-35",
			expectError: false,
		},
		{
			name:        "CPF with mixed formatting",
			cpf:         "111444777-35",
			expected:    "111.444.777-35",
			expectError: false,
		},
		{
			name:        "Too short",
			cpf:         "1114447773",
			expected:    "",
			expectError: true,
		},
		{
			name:        "Too long",
			cpf:         "111444777355",
			expected:    "",
			expectError: true,
		},
		{
			name:        "Empty string",
			cpf:         "",
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := MaskCPF(&MaskCPFInput{CPF: tt.cpf})

			if tt.expectError && err == nil {
				t.Error("MaskCPF() expected error, got nil")
			}

			if !tt.expectError && err != nil {
				t.Errorf("MaskCPF() unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("MaskCPF(%s) = %v, want %v", tt.cpf, result, tt.expected)
			}
		})
	}
}

func TestGenerateCPF(t *testing.T) {
	t.Run("Generate unformatted CPF", func(t *testing.T) {
		cpf := GenerateCPF(&GenerateCPFInput{Formatted: false})

		if len(cpf) != 11 {
			t.Errorf("GenerateCPF() length = %d, want 11", len(cpf))
		}

		if !ValidateCPF(&ValidateCPFInput{CPF: cpf}) {
			t.Errorf("GenerateCPF() generated invalid CPF: %s", cpf)
		}

		for _, c := range cpf {
			if c < '0' || c > '9' {
				t.Errorf("GenerateCPF() contains non-digit character: %c", c)
			}
		}
	})

	t.Run("Generate formatted CPF", func(t *testing.T) {
		cpf := GenerateCPF(&GenerateCPFInput{Formatted: true})

		if len(cpf) != 14 {
			t.Errorf("GenerateCPF() formatted length = %d, want 14", len(cpf))
		}

		if !ValidateCPF(&ValidateCPFInput{CPF: cpf}) {
			t.Errorf("GenerateCPF() generated invalid formatted CPF: %s", cpf)
		}

		if cpf[3] != '.' || cpf[7] != '.' || cpf[11] != '-' {
			t.Errorf("GenerateCPF() incorrect format: %s", cpf)
		}
	})

	t.Run("Generate multiple CPFs - uniqueness", func(t *testing.T) {
		cpfs := make(map[string]bool)
		duplicates := 0

		for range 100 {
			cpf := GenerateCPF(&GenerateCPFInput{Formatted: false})
			if cpfs[cpf] {
				duplicates++
			}
			cpfs[cpf] = true
		}

		if duplicates > 5 {
			t.Errorf("GenerateCPF() generated too many duplicates: %d out of 100", duplicates)
		}
	})
}
