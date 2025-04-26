package user

import (
	"errors"
	"time"
)

type Password struct {
	plaintext *string
	hash      []byte
}

type User struct {
	ID             int64
	Name           string
	DocumentNumber DocumentNumber
	DocumentType   DocumentType
	Email          Email
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

	if password == "" {
		return nil, errors.New("senha não pode ser vazia")
	}

	if documentType == "" {
		return nil, errors.New("tipo de documento não pode ser vazio")
	}

	u := &User{
		Name:           name,
		DocumentNumber: DocumentNumber(documentNumber), // Assuming it's a string type
		DocumentType:   documentType,
		Email:          validatedEmail,
		Password:       password,
		CreatedAt:      time.Now(),
	}
	return u, nil
}
