package cpfvalidator_test

import (
	"testing"

	"github.com/itispx/cpfcnpj/cpfvalidator"
)

func TestInvalidPolluted(t *testing.T) {
	cpf := "012.345.678-99"

	if cpfvalidator.Validate(cpf) {
		t.Error("invalid cpf passing as valid")
	}
}

func TestInvalidClean(t *testing.T) {
	cpf := "01234567899"

	if cpfvalidator.Validate(cpf) {
		t.Error("invalid cpf passing as valid")
	}
}

func TestValidPolluted(t *testing.T) {
	cpf := "197.367.670-27"

	if !cpfvalidator.Validate(cpf) {
		t.Error("invalid cpf passing as valid")
	}
}

func TestValidClean(t *testing.T) {
	cpf := "19736767027"

	if !cpfvalidator.Validate(cpf) {
		t.Error("invalid cpf passing as valid")
	}
}
