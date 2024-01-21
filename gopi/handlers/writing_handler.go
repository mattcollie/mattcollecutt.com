package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gopi/models"
	"gopi/templates"
)

func WritingHandler(c *fiber.Ctx) error {
	socialLinks := []models.SocialLink{
		{Name: "Github", URL: "https://github.com/mattcollie", SVG: "/static/images/github.svg"},
		{Name: "Linkedin", URL: "https://www.linkedin.com/in/matt-collecutt", SVG: "/static/images/linkedin.svg"},
		{Name: "Instagram", URL: "https://www.instagram.com/matthewcollecutt", SVG: "/static/images/instagram.svg"},
	}
	return Render(c, templates.DefaultPage("Writing", socialLinks))
}
