package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/caioedlobo/desafio-picpay-go/internal/application/commands"
	"github.com/caioedlobo/desafio-picpay-go/internal/application/handlers"
	"github.com/caioedlobo/desafio-picpay-go/internal/domain/user"
	"github.com/gofiber/fiber/v2"
)

type HTTPHandler struct {
	commandHandler *handlers.CommandHandler
	//queryHandler   *handlers.QueryHandler
}

func NewHTTPHandler(commandHandler *handlers.CommandHandler) *HTTPHandler {
	return &HTTPHandler{commandHandler: commandHandler}
}

func (h *HTTPHandler) CreateUser(c *fiber.Ctx) error {
	createInput := commands.CreateUserCommand{}
	err := readJSON(c, &createInput)
	if err != nil {
		return handlers.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if _, err := user.ValidEmail(createInput.Email); err != nil {
		return handlers.BadRequestErrorResponse(c, handlers.ErrEmailNotValid.Error())
	}

	if _, err := user.ValidName(createInput.Name); err != nil {
		return handlers.BadRequestErrorResponse(c, handlers.ErrEmailNotValid.Error())
	}

	if _, err := user.ValidPassword(createInput.Name); err != nil {
		return handlers.BadRequestErrorResponse(c, handlers.ErrEmailNotValid.Error())
	}
	if _, err := user.ValidDocumentNumber(createInput.Name); err != nil {
		return handlers.BadRequestErrorResponse(c, handlers.ErrEmailNotValid.Error())
	}
	if _, err := user.ValidDocumentType(createInput.Name); err != nil {
		return handlers.BadRequestErrorResponse(c, handlers.ErrEmailNotValid.Error())
	}

	id, err := h.commandHandler.HandleCreateUser(context.Background(), createInput)
	if err != nil {
		switch {
		case errors.Is(err, handlers.ErrEmailAlreadyExists):
			return handlers.FailedValidationErrorResponse(c, handlers.ErrEmailAlreadyExists.Error())
		default:
			return handlers.ServerErrorResponse(c, err)
		}
	}
	return c.JSON(id)

}

func readJSON(c *fiber.Ctx, dst any) error {
	reqBody := c.Request().Body()
	dec := json.NewDecoder(bytes.NewReader(reqBody))
	dec.DisallowUnknownFields()
	return dec.Decode(dst)
}
