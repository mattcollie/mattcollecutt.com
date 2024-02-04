package middleware

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/mattcollie/mattcollecutt.com/service"
)

func WithAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		isAuthenticated := false
		jwtToken, err := c.Cookie("jwt")
		if err == nil {
			if jwtToken.Value != "" {
				err = service.ValidateAccessToken(jwtToken.Value)
				if err == nil {
					isAuthenticated = true
				}
			}
		}
		ctx := context.WithValue(c.Request().Context(), "isAuthenticated", isAuthenticated)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
