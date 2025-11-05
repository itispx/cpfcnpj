package cpfcnpj

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

type ValidateCPFInput struct {
	CPF string
}

func ValidateCPF(input *ValidateCPFInput) bool {
	cpf := Clean(input.CPF)

	if len(cpf) != 11 ||
		cpf == "00000000000" || cpf == "11111111111" ||
		cpf == "22222222222" || cpf == "33333333333" ||
		cpf == "44444444444" || cpf == "55555555555" ||
		cpf == "66666666666" || cpf == "77777777777" ||
		cpf == "88888888888" || cpf == "99999999999" {
		return false
	}

	sum := 0
	for i := range 9 {
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
	for i := range 10 {
		num, _ := strconv.Atoi(string(cpf[i]))
		sum += num * (11 - i)
	}

	secondDigit := (sum * 10) % 11
	if secondDigit == 10 {
		secondDigit = 0
	}

	return secondDigit == int(cpf[10]-'0')
}

type MaskCPFInput struct {
	CPF string
}

func MaskCPF(input *MaskCPFInput) (string, error) {
	cpf := Clean(input.CPF)
	if len(cpf) != 11 {
		return "", errors.New("cpf deve conter 11 digitos")
	}
	return fmt.Sprintf("%s.%s.%s-%s", cpf[0:3], cpf[3:6], cpf[6:9], cpf[9:11]), nil
}

type GenerateCPFInput struct {
	Formatted bool
}

func GenerateCPF(input *GenerateCPFInput) string {
	cpf := make([]int, 11)
	for i := range 9 {
		cpf[i] = rand.Intn(10)
	}

	sum := 0
	for i := range 9 {
		sum += cpf[i] * (10 - i)
	}
	firstDigit := (sum * 10) % 11
	if firstDigit == 10 {
		firstDigit = 0
	}
	cpf[9] = firstDigit

	sum = 0
	for i := range 10 {
		sum += cpf[i] * (11 - i)
	}
	secondDigit := (sum * 10) % 11
	if secondDigit == 10 {
		secondDigit = 0
	}
	cpf[10] = secondDigit

	result := ""
	for _, digit := range cpf {
		result += fmt.Sprintf("%d", digit)
	}

	if input.Formatted {
		formatted, _ := MaskCPF(&MaskCPFInput{
			CPF: result,
		})

		return formatted
	}

	return result
}
