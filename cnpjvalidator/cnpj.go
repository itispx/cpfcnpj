package cnpjvalidator

import (
	"strconv"

	"github.com/itispx/cpfcnpj/strutil"
)

func Validate(cnpj string) bool {
	cnpj = strutil.ExtractNumbers(cnpj)

	if len(cnpj) != 14 {
		return false
	}

	var sum int
	var weight = []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	for i := 0; i < 12; i++ {
		num, _ := strconv.Atoi(string(cnpj[i]))
		sum += num * weight[i]
	}

	firstDigit := sum % 11
	if firstDigit < 2 {
		firstDigit = 0
	} else {
		firstDigit = 11 - firstDigit
	}

	if firstDigit != int(cnpj[12]-'0') {
		return false
	}

	sum = 0
	weight = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	for i := 0; i < 13; i++ {
		num, _ := strconv.Atoi(string(cnpj[i]))
		sum += num * weight[i]
	}

	secondDigit := sum % 11
	if secondDigit < 2 {
		secondDigit = 0
	} else {
		secondDigit = 11 - secondDigit
	}

	return secondDigit == int(cnpj[13]-'0')
}
