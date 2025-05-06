package user

import (
	"context"
	"github.com/caioedlobo/desafio-picpay-go/internal/domain/user/value_object"
)

type UserRepository interface {
	Save(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id int64) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByDocument(ctx context.Context, documentNumber string, documentType value_object.DocumentType) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int64) error
}
