package data

import (
	"database/sql"
	"errors"
	"github.com/caioedlobo/desafio-picpay-go/validator"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Document string

const (
	CPF  Document = "cpf"
	CNPJ Document = "cnpj"
)

type password struct {
	plaintext *string
	hash      []byte
}

type User struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	DocumentNumber string    `json:"document_number"`
	DocumentType   Document  `json:"document_type"`
	Email          string    `json:"email"`
	Password       password  `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
}

type UserModel struct {
	DB *sql.DB
}

func (m UserModel) Insert(user *User) error {
	query := ` INSERT INTO users (name, document_number, document_type, email, password_hash) 
 			   	VALUES($1, $2, $3, $4, $5)
 				RETURNING id, created_at`

	args := []any{user.Name, user.DocumentNumber, user.DocumentType, user.Email, user.Password.hash}

	err := m.DB.QueryRow(query, args...).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (p *password) Set(plainTextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainTextPassword), 12)
	if err != nil {
		return err
	}
	p.plaintext = &plainTextPassword
	p.hash = hash
	return nil
}

func (p *password) Matches(plainTextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(plainTextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil

		default:
			return false, err
		}
	}
	return true, nil
}

func ValidateUser(v *validator.Validator, user User) {
	ValidateName(v, user.Name)
	ValidateEmail(v, user.Email)
	ValidateDocumentNumber(v, user.DocumentNumber, string(user.DocumentType))
	if user.Password.plaintext != nil {
		ValidatePasswordPlaintext(v, *user.Password.plaintext)
	}
	if user.Password.hash == nil {
		panic("missing password hash for user")
	}

}

func ValidateName(v *validator.Validator, name string) {
	v.Check(name != "", "name", "must be provided")
	v.Check(len(name) <= 500, "name", "must not be more than 500 bytes long")
}

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(validator.Matches(email, validator.EmailRX), "email", "must be a valid email address")
}

func ValidateDocumentNumber(v *validator.Validator, documentNumber string, documentType string) {
	v.Check(documentNumber != "", "documentNumber", "must be provided")
	v.Check(documentType != "", "documentType", "must be provided")
	if documentType == string(CPF) {
		v.Check(validator.Matches(documentNumber, validator.CpfRegex), "documentNumber", "must be a valid "+documentType)
	} else if documentType == string(CNPJ) {
		v.Check(validator.Matches(documentNumber, validator.CnpjRegex), "documentNumber", "must be a valid "+documentType)
	} else {
		v.AddError("documentType", "invalid documentType: "+documentType)
	}
}

func ValidatePasswordPlaintext(v *validator.Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 bytes long")
	v.Check(len(password) <= 72, "password", "must not be more than 72 bytes long")
}
