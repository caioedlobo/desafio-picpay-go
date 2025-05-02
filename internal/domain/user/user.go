package user

import (
	"errors"
	"github.com/bojanz/currency"
	"time"
)

type User struct {
	ID             int64
	Name           string
	DocumentNumber DocumentNumber
	DocumentType   DocumentType
	Email          Email
	Password       Password
	Balance        currency.Amount
	CreatedAt      time.Time
}

func NewUser(name string, documentNumber string, password Password, documentType DocumentType, email string) (*User, error) {
	if name == "" {
		return nil, errors.New("nome não pode ser vazio")
	}

	if documentNumber == "" {
		return nil, errors.New("número do documento não pode ser vazio")
	}

	if password.plaintext == nil {
		return nil, errors.New("senha não pode ser vazia")
	}

	if documentType == "" {
		return nil, errors.New("tipo de documento não pode ser vazio")
	}
	balance, err := currency.NewAmount("0.0", "BRL")
	if err != nil {
		return nil, err
	}

	u := &User{
		Name:           name,
		DocumentNumber: DocumentNumber(documentNumber),
		DocumentType:   documentType,
		Email:          Email(email),
		Password:       password,
		Balance:        balance,
		CreatedAt:      time.Now(),
	}
	return u, nil
}
