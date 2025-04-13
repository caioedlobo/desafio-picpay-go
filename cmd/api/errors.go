package main

import (
	"github.com/gofiber/fiber/v2"
)

func (app *application) errorResponse(c *fiber.Ctx, status int, message any) error {
	c.Status(status)
	err := c.JSON(fiber.Map{"error": message})
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
	}
	return err
}

func (app *application) badRequestResponse(c *fiber.Ctx, err error) error {
	return app.errorResponse(c, fiber.StatusBadRequest, err.Error())
}

func (app *application) failedValidationResponse(c *fiber.Ctx, errors map[string]string) error {
	return app.errorResponse(c, fiber.StatusUnprocessableEntity, errors)
}
