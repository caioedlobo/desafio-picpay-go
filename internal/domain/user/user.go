package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bojanz/currency"
	"github.com/caioedlobo/desafio-picpay-go/internal/domain"
	"github.com/caioedlobo/desafio-picpay-go/internal/domain/event"
	"github.com/caioedlobo/desafio-picpay-go/internal/domain/user/value_object"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID             uuid.UUID
	Name           string
	DocumentNumber value_object.DocumentNumber
	DocumentType   value_object.DocumentType
	Email          value_object.Email
	Password       value_object.Password
	Balance        currency.Amount
	CreatedAt      time.Time
	Aggregate      *domain.Aggregate
}

func NewUser(name string, documentNumber string, password value_object.Password, documentType value_object.DocumentType, email string) (*User, error) {
	if name == "" {
		return nil, errors.New("nome não pode ser vazio")
	}

	if documentNumber == "" {
		return nil, errors.New("número do documento não pode ser vazio")
	}

	if password.GetPlaintext() == nil {
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
		ID:             uuid.New(),
		Name:           name,
		DocumentNumber: value_object.DocumentNumber(documentNumber),
		DocumentType:   documentType,
		Email:          value_object.Email(email),
		Password:       password,
		Balance:        balance,
		CreatedAt:      time.Now(),
	}
	u.Aggregate = domain.NewAggregate(u.ID.String(), u.ApplyEvent)
	return u, nil
}

func (u *User) ApplyEvent(ev *event.Event) {
	eventData := &User{}
	err := json.Unmarshal(ev.Data, &eventData)
	if err != nil {
		return
	}

	switch ev.Type {
	case event.UserCreated:
		u.ID = eventData.ID
		u.Name = eventData.Name
		u.Email = eventData.Email
		u.DocumentType = eventData.DocumentType
		u.DocumentNumber = eventData.DocumentNumber
		u.Password = eventData.Password
	case event.UserNameUpdated:
		u.Name = eventData.Name
	default:
		panic(fmt.Sprintf("unknown event: %s", ev.Type))
	}
}
