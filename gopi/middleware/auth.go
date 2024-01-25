package middleware

import (
	"github.com/gofiber/fiber/v2"
	"gopi/services"
)

func AuthMiddleware(c *fiber.Ctx) error {
	isAuthenticated := false
	jwtToken := c.Cookies("jwt")
	if jwtToken != "" {
		err := services.ValidateAccessToken(jwtToken)
		if err == nil {
			isAuthenticated = true
		}
	}

	c.Locals("isAuthenticated", isAuthenticated)

	return c.Next()
}
