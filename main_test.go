package cpfcnpj_test

import (
	"testing"

	"github.com/itispx/cpfcnpj/cnpjvalidator"
	"github.com/itispx/cpfcnpj/cpfvalidator"
)

func TestCPF_InvalidPolluted(t *testing.T) {
	cpf := "012.345.678-99"

	if cpfvalidator.Validate(cpf) {
		t.Error("invalid cpf passing as valid")
	}
}

func TestCPF_InvalidClean(t *testing.T) {
	cpf := "01234567899"

	if cpfvalidator.Validate(cpf) {
		t.Error("invalid cpf passing as valid")
	}
}

func TestCPF_ValidPolluted(t *testing.T) {
	cpf := "197.367.670-27"

	if !cpfvalidator.Validate(cpf) {
		t.Error("invalid cpf passing as valid")
	}
}

func TestCPF_ValidClean(t *testing.T) {
	cpf := "19736767027"

	if !cpfvalidator.Validate(cpf) {
		t.Error("invalid cpf passing as valid")
	}
}

func TestCNPJ_InvalidPolluted(t *testing.T) {
	cnpj := "01.234.567/8999-99"

	if cnpjvalidator.Validate(cnpj) {
		t.Error("invalid cnpj passing as valid")
	}
}

func TestCNPJ_InvalidClean(t *testing.T) {
	cnpj := "01234567899999"

	if cnpjvalidator.Validate(cnpj) {
		t.Error("invalid cnpj passing as valid")
	}
}

func TestCNPJ_ValidPolluted(t *testing.T) {
	cnpj := "01.009.829/0001-18"

	if !cnpjvalidator.Validate(cnpj) {
		t.Error("invalid cnpj passing as valid")
	}
}

func TestCNPJ_ValidClean(t *testing.T) {
	cnpj := "01009829000118"

	if !cnpjvalidator.Validate(cnpj) {
		t.Error("invalid cnpj passing as valid")
	}
}
