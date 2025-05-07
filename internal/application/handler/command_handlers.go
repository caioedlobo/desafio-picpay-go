package handler

import (
	"context"
	"encoding/json"
	"github.com/caioedlobo/desafio-picpay-go/internal/application/command"
	"github.com/caioedlobo/desafio-picpay-go/internal/domain/event"
	"github.com/caioedlobo/desafio-picpay-go/internal/domain/user"
	"github.com/caioedlobo/desafio-picpay-go/internal/domain/user/dto"
	"github.com/caioedlobo/desafio-picpay-go/internal/domain/user/value_object"
)

type CommandHandler struct {
	userRepo  user.UserRepository
	eventRepo event.EventRepository
}

func NewCommandHandler(userRepo user.UserRepository, eventStore event.EventRepository) *CommandHandler {
	return &CommandHandler{
		userRepo:  userRepo,
		eventRepo: eventStore,
	}
}
func (h *CommandHandler) HandleCreateUser(ctx context.Context, cmd command.CreateUserCommand) (string, error) {
	if existingUser, _ := h.userRepo.FindByEmail(ctx, cmd.Email); existingUser != nil {
		return "", ErrEmailAlreadyExists
	}

	docType, err := value_object.NewDocumentType(cmd.DocumentType)
	if err != nil {
		return "", err
	}

	pass, err := value_object.NewPassword(cmd.Password)
	if err != nil {
		return "", err
	}

	u, err := user.NewUser(cmd.Name, cmd.DocumentNumber, *pass, docType, cmd.Email)
	if err != nil {
		return "", err
	}

	if err = h.userRepo.Save(ctx, u); err != nil {
		return "", err
	}

	createUser, _ := json.Marshal(dto.NewCreateUser(u))
	u.Aggregate.AddEvent(event.UserCreated, createUser)

	if err = h.eventRepo.AppendEvent(ctx, u.Aggregate.Events()); err != nil {
		return "", err
	}

	return "Ok", nil
}
