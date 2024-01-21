package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gopi/templates"
)

func RootHandler(c *fiber.Ctx) error {
	return Render(c, templates.HomePage())
}
