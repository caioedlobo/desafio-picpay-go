package handler

import (
	"context"
	"encoding/json"
	"github.com/caioedlobo/desafio-picpay-go/internal/application/command"
	"github.com/caioedlobo/desafio-picpay-go/internal/domain/event"
	"github.com/caioedlobo/desafio-picpay-go/internal/domain/user"
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
func (h *CommandHandler) HandleCreateUser(ctx context.Context, cmd command.CreateUserCommand) error {
	if existingUser, _ := h.userRepo.FindByEmail(ctx, cmd.Email); existingUser != nil {
		return ErrEmailAlreadyExists
	}

	docType, err := value_object.NewDocumentType(cmd.DocumentType)
	if err != nil {
		return err
	}

	pass, err := value_object.NewPassword(cmd.Password)
	if err != nil {
		return err
	}

	u, err := user.NewUser(cmd.Name, cmd.DocumentNumber, *pass, docType, cmd.Email)
	if err != nil {
		return err
	}

	if err = h.userRepo.Save(ctx, u); err != nil {
		return err
	}

	dto, _ := json.Marshal(cmd)
	u.Aggregate.AddEvent(event.UserCreated, dto)

	if err = h.eventRepo.AppendEvent(ctx, u.Aggregate.Events()); err != nil {
		return err
	}

	return nil
}

func (h *CommandHandler) HandleUpdateUserName(ctx context.Context, cmd command.UpdateUserNameCommand) error {
	if existingUser, _ := h.userRepo.FindByID(ctx, cmd.ID); existingUser == nil {
		return ErrEmailAlreadyExists
	}

	err := h.userRepo.UpdateName(ctx, cmd)
	if err != nil {
		return err
	}

	agg, err := h.eventRepo.Get(ctx, cmd.ID)
	if err != nil {
		return err
	}

	dto, err := json.Marshal(cmd)
	if err != nil {
		return err
	}

	agg.AddEvent(event.UserNameUpdated, dto)

	err = h.eventRepo.AppendEvent(ctx, agg.Events())
	if err != nil {
		return err
	}

	return nil
}
