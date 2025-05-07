package value_object

import (
	"errors"
)

const (
	CPF  DocumentType = "cpf"
	CNPJ DocumentType = "cnpj"
)

type DocumentType string

func NewDocumentType(value string) (DocumentType, error) {
	if ok, err := ValidDocumentType(value); !ok {
		return "", err
	}
	return DocumentType(value), nil
}

func ValidDocumentType(value string) (bool, error) {

	if value == "" {
		return false, errors.New("tipo de documento não pode estar vazio")
	}
	if value != string(CPF) && value != string(CNPJ) {
		return false, errors.New("tipo de documento inválido")
	}
	return true, nil
}
