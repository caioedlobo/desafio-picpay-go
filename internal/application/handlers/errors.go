package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrUserNotFound       = errors.New("usuário não encontrado")
	ErrEmailAlreadyExists = errors.New("email já está em uso")
	ErrInternalServer     = errors.New("ocorreu um erro inesperado")
	ErrEmailNotValid      = errors.New("email não é válido")
)

func ErrorResponse(c *fiber.Ctx, status int, message any) error {
	c.Status(status)
	err := c.JSON(fiber.Map{"error": message})
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
	}
	return err
}

func ServerErrorResponse(c *fiber.Ctx, err error) error {
	return ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
}

func BadRequestErrorResponse(c *fiber.Ctx, message string) error {
	return ErrorResponse(c, fiber.StatusBadRequest, message)
}

func FailedValidationErrorResponse(c *fiber.Ctx, message string) error {
	return ErrorResponse(c, fiber.StatusUnprocessableEntity, message)
}
