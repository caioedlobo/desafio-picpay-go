package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func (app *application) serve() error {
	srv := fiber.New()

	srv.Use(logger.New())

	srv.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	//srv.Use(middleware.Logger())
	//router.SetupRoutes(srv)

	err := srv.Listen(":3000")
	if err != nil {
		return err
	}
	return nil
}
