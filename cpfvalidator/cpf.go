package cpfvalidator

import (
	"strconv"

	"github.com/itispx/cpfcnpj/strutil"
)

func Validate(cpf string) bool {
	cpf = strutil.ExtractNumbers(cpf)

	if len(cpf) != 11 ||
		cpf == "00000000000" || cpf == "11111111111" ||
		cpf == "22222222222" || cpf == "33333333333" ||
		cpf == "44444444444" || cpf == "55555555555" ||
		cpf == "66666666666" || cpf == "77777777777" ||
		cpf == "88888888888" || cpf == "99999999999" {
		return false
	}

	sum := 0
	for i := 0; i < 9; i++ {
		num, _ := strconv.Atoi(string(cpf[i]))
		sum += num * (10 - i)
	}

	firstDigit := (sum * 10) % 11
	if firstDigit == 10 {
		firstDigit = 0
	}

	if firstDigit != int(cpf[9]-'0') {
		return false
	}

	sum = 0
	for i := 0; i < 10; i++ {
		num, _ := strconv.Atoi(string(cpf[i]))
		sum += num * (11 - i)
	}

	secondDigit := (sum * 10) % 11
	if secondDigit == 10 {
		secondDigit = 0
	}

	return secondDigit == int(cpf[10]-'0')
}
