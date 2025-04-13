package main

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func (app *application) serve() error {
	srv := fiber.New()

	srv.Use(logger.New())

	srv.Post("api/v1/users", app.registerUserHandler)

	//srv.Use(middleware.Logger())
	//router.SetupRoutes(srv)

	err := srv.Listen(":3000")
	if err != nil {
		return err
	}
	return nil
}
func (app *application) strictBodyParser(c *fiber.Ctx, dst any) error {
	body := c.Request().Body()
	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.DisallowUnknownFields()
	return decoder.Decode(dst)
}
