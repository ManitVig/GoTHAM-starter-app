package helpers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/manitvig/gotham-starter-app/internal/handlers"
)

func SetupHandlers(app *fiber.App) {
	app.Static("/", "./public/")
	
	app.Get("/", handlers.IndexPage)
	app.Post("/increment", handlers.Increment)
}
