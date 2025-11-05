package cpfcnpj

import (
	"testing"
)

func TestValidateCNPJ(t *testing.T) {
	tests := []struct {
		name     string
		cnpj     string
		expected bool
	}{
		{
			name:     "Valid CNPJ 1",
			cnpj:     "11222333000181",
			expected: true,
		},
		{
			name:     "Valid CNPJ 2",
			cnpj:     "34028316000103",
			expected: true,
		},
		{
			name:     "Valid CNPJ with formatting",
			cnpj:     "11.222.333/0001-81",
			expected: true,
		},
		{
			name:     "Valid CNPJ - known valid",
			cnpj:     "00000000000191",
			expected: true,
		},
		{
			name:     "Too short",
			cnpj:     "1122233300018",
			expected: false,
		},
		{
			name:     "Too long",
			cnpj:     "112223330001811",
			expected: false,
		},
		{
			name:     "Empty string",
			cnpj:     "",
			expected: false,
		},
		{
			name:     "Wrong first check digit",
			cnpj:     "11222333000191",
			expected: false,
		},
		{
			name:     "Wrong second check digit",
			cnpj:     "11222333000182",
			expected: false,
		},
		{
			name:     "Both check digits wrong",
			cnpj:     "11222333000100",
			expected: false,
		},
		{
			name:     "All zeros",
			cnpj:     "00000000000000",
			expected: false,
		},
		{
			name:     "All ones",
			cnpj:     "11111111111111",
			expected: false,
		},
		{
			name:     "All twos",
			cnpj:     "22222222222222",
			expected: false,
		},
		{
			name:     "All threes",
			cnpj:     "33333333333333",
			expected: false,
		},
		{
			name:     "All fours",
			cnpj:     "44444444444444",
			expected: false,
		},
		{
			name:     "All fives",
			cnpj:     "55555555555555",
			expected: false,
		},
		{
			name:     "All sixes",
			cnpj:     "66666666666666",
			expected: false,
		},
		{
			name:     "All sevens",
			cnpj:     "77777777777777",
			expected: false,
		},
		{
			name:     "All eights",
			cnpj:     "88888888888888",
			expected: false,
		},
		{
			name:     "All nines",
			cnpj:     "99999999999999",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateCNPJ(&ValidateCNPJInput{CNPJ: tt.cnpj})
			if result != tt.expected {
				t.Errorf("ValidateCNPJ(%s) = %v, want %v", tt.cnpj, result, tt.expected)
			}
		})
	}
}

func TestMaskCNPJ(t *testing.T) {
	tests := []struct {
		name        string
		cnpj        string
		expected    string
		expectError bool
	}{
		{
			name:        "Valid CNPJ",
			cnpj:        "11222333000181",
			expected:    "11.222.333/0001-81",
			expectError: false,
		},
		{
			name:        "CNPJ already formatted",
			cnpj:        "11.222.333/0001-81",
			expected:    "11.222.333/0001-81",
			expectError: false,
		},
		{
			name:        "CNPJ with mixed formatting",
			cnpj:        "11222333/0001-81",
			expected:    "11.222.333/0001-81",
			expectError: false,
		},
		{
			name:        "Too short",
			cnpj:        "1122233300018",
			expected:    "",
			expectError: true,
		},
		{
			name:        "Too long",
			cnpj:        "112223330001811",
			expected:    "",
			expectError: true,
		},
		{
			name:        "Empty string",
			cnpj:        "",
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := MaskCNPJ(&MaskCNPJInput{CNPJ: tt.cnpj})

			if tt.expectError && err == nil {
				t.Error("MaskCNPJ() expected error, got nil")
			}

			if !tt.expectError && err != nil {
				t.Errorf("MaskCNPJ() unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("MaskCNPJ(%s) = %v, want %v", tt.cnpj, result, tt.expected)
			}
		})
	}
}

func TestGenerateCNPJ(t *testing.T) {
	t.Run("Generate unformatted CNPJ", func(t *testing.T) {
		cnpj := GenerateCNPJ(&GenerateCNPJInput{Formatted: false})

		if len(cnpj) != 14 {
			t.Errorf("GenerateCNPJ() length = %d, want 14", len(cnpj))
		}

		if !ValidateCNPJ(&ValidateCNPJInput{CNPJ: cnpj}) {
			t.Errorf("GenerateCNPJ() generated invalid CNPJ: %s", cnpj)
		}

		for _, c := range cnpj {
			if c < '0' || c > '9' {
				t.Errorf("GenerateCNPJ() contains non-digit character: %c", c)
			}
		}

		if cnpj[8:12] != "0001" {
			t.Errorf("GenerateCNPJ() positions 8-11 = %s, want 0001", cnpj[8:12])
		}
	})

	t.Run("Generate formatted CNPJ", func(t *testing.T) {
		cnpj := GenerateCNPJ(&GenerateCNPJInput{Formatted: true})

		if len(cnpj) != 18 {
			t.Errorf("GenerateCNPJ() formatted length = %d, want 18", len(cnpj))
		}

		if !ValidateCNPJ(&ValidateCNPJInput{CNPJ: cnpj}) {
			t.Errorf("GenerateCNPJ() generated invalid formatted CNPJ: %s", cnpj)
		}

		if cnpj[2] != '.' || cnpj[6] != '.' || cnpj[10] != '/' || cnpj[15] != '-' {
			t.Errorf("GenerateCNPJ() incorrect format: %s", cnpj)
		}
	})

	t.Run("Generate multiple CNPJs - uniqueness", func(t *testing.T) {
		cnpjs := make(map[string]bool)
		duplicates := 0

		for range 100 {
			cnpj := GenerateCNPJ(&GenerateCNPJInput{Formatted: false})
			if cnpjs[cnpj] {
				duplicates++
			}
			cnpjs[cnpj] = true
		}

		if duplicates > 5 {
			t.Errorf("GenerateCNPJ() generated too many duplicates: %d out of 100", duplicates)
		}
	})

	t.Run("Validate CNPJ check digit calculation", func(t *testing.T) {
		for range 50 {
			cnpj := GenerateCNPJ(&GenerateCNPJInput{Formatted: false})
			if !ValidateCNPJ(&ValidateCNPJInput{CNPJ: cnpj}) {
				t.Errorf("Generated CNPJ failed validation: %s", cnpj)
			}
		}
	})
}
