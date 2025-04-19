package main

import "github.com/gofiber/fiber/v2"

func (app *application) healthcheckHandler(c *fiber.Ctx) error {
	env := fiber.Map{
		"status": "avaialble",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}
	return c.JSON(env)
}
