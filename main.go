package cpfcnpj

import (
	"github.com/itispx/cpfcnpj/cnpjvalidator"
	"github.com/itispx/cpfcnpj/cpfvalidator"
	"github.com/itispx/cpfcnpj/strutil"
)

func Validate(v string) (bool, string) {
	v = strutil.ExtractNumbers(v)

	if len(v) == 11 {
		return cpfvalidator.Validate(v), "cpf"
	} else if len(v) == 14 {
		return cnpjvalidator.Validate(v), "cnpj"
	}

	return false, ""
}
