package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gopi/templates"
)

func WritingHandler(c *fiber.Ctx) error {
	return Render(c, templates.DefaultPage("Writing"))
}
