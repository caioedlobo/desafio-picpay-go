package dto

import (
	"github.com/caioedlobo/desafio-picpay-go/internal/domain/user"
	"github.com/caioedlobo/desafio-picpay-go/internal/domain/user/value_object"
	"github.com/google/uuid"
	"time"
)

type CreateUser struct {
	ID             uuid.UUID                   `json:"id"`
	Name           string                      `json:"name"`
	DocumentNumber value_object.DocumentNumber `json:"document_number"`
	DocumentType   value_object.DocumentType   `json:"document_type"`
	Email          value_object.Email          `json:"email"`
	CreatedAt      time.Time                   `json:"created_at"`
}

func NewCreateUser(u *user.User) CreateUser {
	return CreateUser{
		ID:             u.ID,
		Name:           u.Name,
		DocumentNumber: u.DocumentNumber,
		DocumentType:   u.DocumentType,
		Email:          u.Email,
		CreatedAt:      u.CreatedAt,
	}
}
