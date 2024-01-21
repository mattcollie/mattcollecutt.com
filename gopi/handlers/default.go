package handlers

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"gopi/templates"
)

func Render(c *fiber.Ctx, component templ.Component, options ...func(handler *templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}

	return adaptor.HTTPHandler(componentHandler)(c)
}

func NotFoundHandler(c *fiber.Ctx) error {
	c.Status(fiber.StatusNotFound)
	return Render(c, templates.NotFoundPage())
}
