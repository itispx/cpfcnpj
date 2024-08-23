package cnpjvalidator_test

import (
	"testing"

	"github.com/itispx/cpfcnpj/cnpjvalidator"
)

func TestInvalidPolluted(t *testing.T) {
	cnpj := "01.234.567/8999-99"

	if cnpjvalidator.Validate(cnpj) {
		t.Error("invalid cnpj passing as valid")
	}
}

func TestInvalidClean(t *testing.T) {
	cnpj := "01234567899999"

	if cnpjvalidator.Validate(cnpj) {
		t.Error("invalid cnpj passing as valid")
	}
}

func TestValidPolluted(t *testing.T) {
	cnpj := "01.009.829/0001-18"

	if !cnpjvalidator.Validate(cnpj) {
		t.Error("invalid cnpj passing as valid")
	}
}

func TestValidClean(t *testing.T) {
	cnpj := "01009829000118"

	if !cnpjvalidator.Validate(cnpj) {
		t.Error("invalid cnpj passing as valid")
	}
}
