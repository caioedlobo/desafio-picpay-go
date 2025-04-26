package user

import (
	"errors"
	"regexp"
)

type Email string

func (e Email) New(value string) (Email, error) {
	if ok, err := ValidEmail(value); !ok {
		return "", err
	}
	return Email(value), nil
}

func ValidEmail(value string) (bool, error) {
	var EmailRX = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if value == "" {
		return false, errors.New("email não pode ser vazio")
	}
	if ok := EmailRX.MatchString(value); !ok {
		return false, errors.New("email deve ser válido")
	}
	return true, nil
}
