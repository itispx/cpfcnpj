package cpfcnpj

import "fmt"

type ValidateInput struct {
	CPFCNPJ string
}

func Validate(input *ValidateInput) (bool, string) {
	cpfcnpj := Clean(input.CPFCNPJ)

	if len(cpfcnpj) == 11 {
		return ValidateCPF(&ValidateCPFInput{
			CPF: cpfcnpj,
		}), "cpf"
	} else if len(cpfcnpj) == 14 {
		return ValidateCNPJ(&ValidateCNPJInput{
			CNPJ: cpfcnpj,
		}), "cnpj"
	}

	return false, ""
}

type MaskInput struct {
	CPFCNPJ string
}

func Mask(input *MaskInput) (string, error) {
	cpfcnpj := Clean(input.CPFCNPJ)

	switch len(cpfcnpj) {
	case 11:
		return MaskCPF(&MaskCPFInput{
			CPF: cpfcnpj,
		})
	case 14:
		return MaskCNPJ(&MaskCNPJInput{
			CNPJ: cpfcnpj,
		})
	default:
		return cpfcnpj, fmt.Errorf("documento inválido: esperado 11 dígitos (CPF) ou 14 dígitos (CNPJ), mas recebeu %d", len(cpfcnpj))
	}
}
