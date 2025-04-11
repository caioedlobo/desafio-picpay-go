package main

import (
	"github.com/gofiber/fiber/v2"
)

func (app *application) serve() error {
	srv := fiber.New()

	srv.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	err := srv.Listen(":3000")
	if err != nil {
		return err
	}
	return nil
}
