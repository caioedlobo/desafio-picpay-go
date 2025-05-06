package value_object

import (
	"errors"
	"regexp"
)

type DocumentNumber string

func (e DocumentNumber) New(value string) (DocumentNumber, error) {
	if ok, err := ValidDocumentNumber(value); !ok {
		return "", err
	}
	return DocumentNumber(value), nil
}

func ValidDocumentNumber(value string) (bool, error) {
	var CpfRegex = regexp.MustCompile(`^\d{11}$`)
	var CnpjRegex = regexp.MustCompile(`^\d{14}$`)
	if value == "" {
		return false, errors.New("CPF/CNPJ não pode estar vazio")
	}
	if ok := CpfRegex.MatchString(value) || CnpjRegex.MatchString(value); !ok {
		return false, errors.New("deve ser um CPF/CNPJ válido")
	}
	return true, nil
}
