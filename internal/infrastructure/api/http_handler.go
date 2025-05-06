package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/caioedlobo/desafio-picpay-go/internal/application/command"
	"github.com/caioedlobo/desafio-picpay-go/internal/application/handlers"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type HTTPHandler struct {
	commandHandler *handlers.CommandHandler
	queryHandler   *handlers.QueryHandler
	validator      *validator.Validate
}

func NewHTTPHandler(commandHandler *handlers.CommandHandler, queryHandler *handlers.QueryHandler, validator *validator.Validate) *HTTPHandler {
	return &HTTPHandler{
		commandHandler: commandHandler,
		queryHandler:   queryHandler,
		validator:      validator}
}

func (h *HTTPHandler) CreateUser(c *fiber.Ctx) error {
	createInput := command.CreateUserCommand{}
	err := readJSON(c, &createInput)
	if err != nil {
		return handlers.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	err = h.validator.Struct(createInput)
	if err != nil {
		return handlers.BadRequestErrorResponse(c, err.Error())
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
