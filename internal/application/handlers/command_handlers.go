package handlers

import (
	"context"
	"encoding/json"
	"github.com/caioedlobo/desafio-picpay-go/internal/application/commands"
	"github.com/caioedlobo/desafio-picpay-go/internal/domain/event"
	"github.com/caioedlobo/desafio-picpay-go/internal/domain/user"
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
func (h *CommandHandler) HandleCreateUser(ctx context.Context, cmd commands.CreateUserCommand) (int64, error) {
	if existingUser, _ := h.userRepo.FindByEmail(ctx, cmd.Email); existingUser != nil {
		return 0, ErrEmailAlreadyExists
	}

	docType := user.DocumentType(cmd.DocumentType)

	pass, err := user.NewPassword(cmd.Password)
	if err != nil {
		return 0, err
	}

	u, err := user.NewUser(cmd.Name, cmd.DocumentNumber, *pass, docType, cmd.Email)
	if err != nil {
		return 0, err
	}

	if err = h.userRepo.Save(ctx, u); err != nil {
		return 0, err
	}

	// Criar evento
	userData, _ := json.Marshal(map[string]any{
		"id":              u.ID,
		"name":            u.Name,
		"document_number": u.DocumentNumber,
		"document_type":   u.DocumentType,
		"email":           u.Email,
		"created_at":      u.CreatedAt,
	})

	u.Aggregate.AddEvent(event.UserCreated, userData)

	if err = h.eventRepo.AppendEvent(ctx, u.Aggregate.Events()); err != nil {
		return 0, err
	}

	return u.ID, nil
}
