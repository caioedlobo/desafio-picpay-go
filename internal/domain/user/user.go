package user

import (
	"errors"
	"regexp"
	"time"
)

const (
	CPF  DocumentType = "cpf"
	CNPJ DocumentType = "cnpj"
)

var (
	EmailRX   = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	CpfRegex  = regexp.MustCompile(`^\d{11}$`)
	CnpjRegex = regexp.MustCompile(`^\d{14}$`)
)

type DocumentType string

type Password struct {
	plaintext *string
	hash      []byte
}

type User struct {
	ID             int64
	Name           string
	DocumentNumber string
	DocumentType   DocumentType
	Email          string
	Password       string
	CreatedAt      time.Time
}

func NewUser(name string, documentNumber string, password string, documentType DocumentType, email string) (*User, error) {
	if name == "" {
		return nil, errors.New("nome não pode ser vazio")
	}

	if documentNumber == "" {
		return nil, errors.New("número do documento não pode ser vazio")
	}

	if email == "" {
		return nil, errors.New("email não pode ser vazio")
	}

	if password == "" {
		return nil, errors.New("senha não pode ser vazia")
	}

	if documentType == "" {
		return nil, errors.New("tipo de documento não pode ser vazio")
	}

	u := &User{
		Name:           name,
		DocumentNumber: documentNumber,
		DocumentType:   documentType,
		Email:          email,
		Password:       password,
		CreatedAt:      time.Now(),
	}
	return u, nil
}
