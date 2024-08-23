package cpfcnpj_test

import (
	"testing"

	"github.com/itispx/cpfcnpj"
)

func TestCPF_InvalidPolluted(t *testing.T) {
	cpf := "012.345.678-99"

	isValid, _ := cpfcnpj.Validate(cpf)

	if isValid {
		t.Error("invalid cpf passing as valid")
	}
}

func TestCPF_InvalidClean(t *testing.T) {
	cpf := "01234567899"

	isValid, _ := cpfcnpj.Validate(cpf)

	if isValid {
		t.Error("invalid cpf passing as valid")
	}
}

func TestCPF_ValidPolluted(t *testing.T) {
	cpf := "197.367.670-27"

	isValid, _ := cpfcnpj.Validate(cpf)

	if !isValid {
		t.Error("invalid cpf passing as valid")
	}
}

func TestCPF_ValidClean(t *testing.T) {
	cpf := "19736767027"

	isValid, _ := cpfcnpj.Validate(cpf)

	if !isValid {
		t.Error("invalid cpf passing as valid")
	}
}

func TestCPF_DocTypePolluted(t *testing.T) {
	cpf := "197.367.670-27"

	_, docType := cpfcnpj.Validate(cpf)

	if docType != "cpf" {
		t.Error("invalid doc type")
	}
}

func TestCPF_DocTypeClean(t *testing.T) {
	cpf := "19736767027"

	_, docType := cpfcnpj.Validate(cpf)

	if docType != "cpf" {
		t.Error("invalid doc type")
	}
}

func TestCNPJ_InvalidPolluted(t *testing.T) {
	cnpj := "01.234.567/8999-99"

	isValid, _ := cpfcnpj.Validate(cnpj)

	if isValid {
		t.Error("invalid cnpj passing as valid")
	}
}

func TestCNPJ_InvalidClean(t *testing.T) {
	cnpj := "01234567899999"

	isValid, _ := cpfcnpj.Validate(cnpj)

	if isValid {
		t.Error("invalid cnpj passing as valid")
	}
}

func TestCNPJ_ValidPolluted(t *testing.T) {
	cnpj := "01.009.829/0001-18"

	isValid, _ := cpfcnpj.Validate(cnpj)

	if !isValid {
		t.Error("invalid cnpj passing as valid")
	}
}

func TestCNPJ_ValidClean(t *testing.T) {
	cnpj := "01009829000118"

	isValid, _ := cpfcnpj.Validate(cnpj)

	if !isValid {
		t.Error("invalid cnpj passing as valid")
	}
}

func TestCNPJ_DocTypePolluted(t *testing.T) {
	cnpj := "01.009.829/0001-18"

	_, docType := cpfcnpj.Validate(cnpj)

	if docType != "cnpj" {
		t.Error("invalid doc type")
	}
}

func TestCNPJ_DocTypeClean(t *testing.T) {
	cnpj := "01009829000118"

	_, docType := cpfcnpj.Validate(cnpj)

	if docType != "cnpj" {
		t.Error("invalid doc type")
	}
}
