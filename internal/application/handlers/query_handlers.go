package handlers

import "github.com/caioedlobo/desafio-picpay-go/internal/domain/user"

type QueryHandler struct {
	userRepo user.UserRepository
}

func NewQueryHandler(userRepo user.UserRepository) *QueryHandler {
	return &QueryHandler{
		userRepo: userRepo,
	}
}
