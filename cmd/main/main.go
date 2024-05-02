package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/manitvig/gotham-starter-app/internal/helpers"
)

func main() {
	app := fiber.New()
	helpers.SetupHandlers(app)
	app.Listen(":8080")
}
