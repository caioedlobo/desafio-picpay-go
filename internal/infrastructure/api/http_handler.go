package api

import (
	"github.com/caioedlobo/desafio-picpay-go/internal/application/handlers"
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
	return c.JSON("Hello World")
}
