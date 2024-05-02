package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/manitvig/gotham-starter-app/internal/helpers/renderer"
	"github.com/manitvig/gotham-starter-app/templates/components"
	"github.com/manitvig/gotham-starter-app/templates/views"
)

var count uint = 0;

func IndexPage(c fiber.Ctx) error {
	return renderer.RenderTempl(c, views.Index(count))
}

func Increment(c fiber.Ctx) error {
	count += 1
	return renderer.RenderTempl(c, components.Counter(count))
}
