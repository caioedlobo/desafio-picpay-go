package handlers

import "errors"

var (
	ErrUserNotFound       = errors.New("usuário não encontrado")
	ErrEmailAlreadyExists = errors.New("email já está em uso")
)
