package main

import (
	"github.com/caioedlobo/desafio-picpay-go/cmd/internal/data"
	"github.com/caioedlobo/desafio-picpay-go/validator"
	"github.com/gofiber/fiber/v2"
)

func (app *application) registerUserHandler(c *fiber.Ctx) error {
	var input struct {
		Name           string        `json:"name"`
		Email          string        `json:"email"`
		DocumentNumber string        `json:"documentNumber"`
		DocumentType   data.Document `json:"documentType"`
		Password       string        `json:"password"`
	}

	err := app.strictBodyParser(c, &input)
	if err != nil {
		return app.badRequestResponse(c, err)
	}

	user := &data.User{
		Name:           input.Name,
		Email:          input.Email,
		DocumentNumber: input.DocumentNumber,
		DocumentType:   input.DocumentType,
	}

	err = user.Password.Set(input.Password)
	if err != nil {
		return app.serverErrorResponse(c, err)
	}

	v := validator.New()
	data.ValidateUser(v, *user)
	if valid := v.Valid(); !valid {
		return app.failedValidationResponse(c, v.Errors)
	}
	app.logger.Debug().Msg("Successful validation")

	err = app.models.Users.Insert(user)
	if err != nil {
		return app.serverErrorResponse(c, err)
	}

	return c.SendStatus(fiber.StatusCreated)
}
