package handlers

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"gopi/mtx"
	"gopi/templates"
	"net/http"
)

func Render(c *fiber.Ctx, component templ.Component, options ...func(handler *templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}

	// This is potentially very bad... but it works
	return adaptor.HTTPHandler(ContextAwareHandler(c, componentHandler))(c)
}

func ContextAwareHandler(c *fiber.Ctx, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isAuthenticated := c.Locals("isAuthenticated").(bool)
		ctx := mtx.SetIsAuthenticated(r.Context(), isAuthenticated)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func NotFoundHandler(c *fiber.Ctx) error {
	c.Status(fiber.StatusNotFound)
	return Render(c, templates.NotFoundPage())
}
