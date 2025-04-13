package main

import (
	"fmt"
	"github.com/caioedlobo/desafio-picpay-go/cmd/internal/data"
	"github.com/caioedlobo/desafio-picpay-go/validator"
	"github.com/gofiber/fiber/v2"
)

func (app *application) registerUserHandler(c *fiber.Ctx) error {
	var input struct {
		Name           string `json:"name"`
		Email          string `json:"email"`
		DocumentNumber string `json:"documentNumber"`
		Password       string `json:"password"`
	}

	err := app.strictBodyParser(c, &input)
	if err != nil {
		return app.badRequestResponse(c, err)
	}
	v := validator.New()
	data.ValidateEmail(v, input.Email)

	if valid := v.Valid(); !valid {
		return app.failedValidationResponse(c, v.Errors)
	}

	user := &data.User{
		Name:           input.Name,
		Email:          input.Email,
		DocumentNumber: input.DocumentNumber,
		Password:       input.Password,
	}
	fmt.Println(user)
	return c.SendStatus(fiber.StatusCreated)
}
