package handlers

import (
	"context"
	"encoding/json"
	"github.com/caioedlobo/desafio-picpay-go/internal/application/commands"
	"github.com/caioedlobo/desafio-picpay-go/internal/domain/event"
	"github.com/caioedlobo/desafio-picpay-go/internal/domain/user"
	"github.com/google/uuid"
	"strconv"
)

type CommandHandler struct {
	userRepo   user.UserRepository
	eventStore event.EventRepository
}

func NewCommandHandler(userRepo user.UserRepository, eventStore event.EventRepository) *CommandHandler {
	return &CommandHandler{
		userRepo:   userRepo,
		eventStore: eventStore,
	}
}
func (h *CommandHandler) HandleCreateUser(ctx context.Context, cmd commands.CreateUserCommand) (int64, error) {
	if existingUser, _ := h.userRepo.FindByEmail(ctx, cmd.Email); existingUser != nil {
		return 0, ErrEmailAlreadyExists
	}

	docType := user.DocumentType(cmd.DocumentType)
	user, err := user.NewUser(cmd.Name, cmd.DocumentNumber, cmd.Password, docType, cmd.Email)
	if err != nil {
		return 0, err
	}

	if err = h.userRepo.Save(ctx, user); err != nil {
		return 0, err
	}

	// Criar evento
	userData, _ := json.Marshal(map[string]any{
		"id":              user.ID,
		"name":            user.Name,
		"document_number": user.DocumentNumber,
		"document_type":   user.DocumentType,
		"email":           user.Email,
		"created_at":      user.CreatedAt,
	})

	ev := &event.Event{
		ID:          uuid.New().String(),
		Type:        event.UserCreated,
		Data:        userData,
		Timestamp:   user.CreatedAt,
		Version:     1,
		AggregateID: strconv.FormatInt(user.ID, 10),
	}
	// TODO:
	// Salvar evento
	if err = h.eventStore.AppendEvent(ctx, ev); err != nil {
		return 0, err
	}

	return user.ID, nil
}
