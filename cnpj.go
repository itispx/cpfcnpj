package cpfcnpj

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

type ValidateCNPJInput struct {
	CNPJ string
}

func ValidateCNPJ(input *ValidateCNPJInput) bool {
	cnpj := Clean(input.CNPJ)

	if len(cnpj) != 14 ||
		cnpj == "00000000000000" || cnpj == "11111111111111" ||
		cnpj == "22222222222222" || cnpj == "33333333333333" ||
		cnpj == "44444444444444" || cnpj == "55555555555555" ||
		cnpj == "66666666666666" || cnpj == "77777777777777" ||
		cnpj == "88888888888888" || cnpj == "99999999999999" {
		return false
	}

	var sum int
	var weight = []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	for i := range 12 {
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

	for i := range 13 {
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

type MaskCNPJInput struct {
	CNPJ string
}

func MaskCNPJ(input *MaskCNPJInput) (string, error) {
	cnpj := Clean(input.CNPJ)
	if len(cnpj) != 14 {
		return "", errors.New("cnpj deve conter 14 digitos")
	}
	return fmt.Sprintf("%s.%s.%s/%s-%s", cnpj[0:2], cnpj[2:5], cnpj[5:8], cnpj[8:12], cnpj[12:14]), nil
}

type GenerateCNPJInput struct {
	Formatted bool
}

func GenerateCNPJ(input *GenerateCNPJInput) string {
	cnpj := make([]int, 14)
	for i := range 8 {
		cnpj[i] = rand.Intn(10)
	}

	cnpj[8] = 0
	cnpj[9] = 0
	cnpj[10] = 0
	cnpj[11] = 1

	weights1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	sum := 0
	for i := range 12 {
		sum += cnpj[i] * weights1[i]
	}
	firstDigit := sum % 11
	if firstDigit < 2 {
		firstDigit = 0
	} else {
		firstDigit = 11 - firstDigit
	}
	cnpj[12] = firstDigit

	weights2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	sum = 0
	for i := range 13 {
		sum += cnpj[i] * weights2[i]
	}
	secondDigit := sum % 11
	if secondDigit < 2 {
		secondDigit = 0
	} else {
		secondDigit = 11 - secondDigit
	}
	cnpj[13] = secondDigit

	result := ""
	for _, digit := range cnpj {
		result += fmt.Sprintf("%d", digit)
	}

	if input.Formatted {
		formatted, _ := MaskCNPJ(&MaskCNPJInput{
			CNPJ: result,
		})

		return formatted
	}

	return result
}
