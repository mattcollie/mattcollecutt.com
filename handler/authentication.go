package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mattcollie/mattcollecutt.com/config"
	"github.com/mattcollie/mattcollecutt.com/service"
	"github.com/mattcollie/mattcollecutt.com/view/component"
	"net/http"
	"time"
)

type AuthHandler struct{}

func (h AuthHandler) HandleAuthLogin(c echo.Context) error {
	loginURL := fmt.Sprintf(
		"%s/realms/%s/protocol/openid-connect/auth?client_id=%s&redirect_uri=%s&response_type=code&scope=openid&kc_idp_hint=github",
		config.Cfg.Authentication.KeycloakURL,
		config.Cfg.Authentication.Realm,
		config.Cfg.Authentication.ClientID,
		config.Cfg.Authentication.RedirectURI)

	if c.Request().Header.Get("HX-Request") == "true" {
		c.Response().Header().Set("HX-Redirect", loginURL)
		return c.NoContent(http.StatusOK)
	}

	return c.Redirect(http.StatusTemporaryRedirect, loginURL)
}

func (h AuthHandler) HandleAuthLogout(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})
	c.SetCookie(&http.Cookie{
		Name:     "refresh_jwt",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})
	return render(c, component.LoginButton())
}

func (h AuthHandler) HandleAuthCallback(c echo.Context) error {
	code := c.QueryParam("code")

	token, err := service.ExchangeToken(code)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error exchanging code for token: "+err.Error())
	}

	err = service.ValidateAccessToken(token.AccessToken)
	if err != nil {
		fmt.Printf("Error decoding access token: %v\n", err)
		return c.String(http.StatusUnauthorized, "Error decoding access token: "+err.Error())
	}

	c.SetCookie(&http.Cookie{
		Name:     "jwt",
		Value:    token.AccessToken,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(24 * time.Hour),
	})

	c.SetCookie(&http.Cookie{
		Name:     "refresh_jwt",
		Value:    token.RefreshToken,
		Expires:  time.Now().Add(24 * 7 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

func (h AuthHandler) HandleAuthRefresh(c echo.Context) error {
	refreshTokenCookie, err := c.Cookie("refresh_jwt")
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, "refresh_jwt cookie not found")
	}

	newToken, err := service.GetRefresh(refreshTokenCookie.Value)
	if err != nil {
		return c.String(http.StatusUnauthorized, "Failed to refresh token")
	}

	c.SetCookie(&http.Cookie{
		Name:     "jwt",
		Value:    newToken.AccessToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	c.SetCookie(&http.Cookie{
		Name:     "refresh_jwt",
		Value:    newToken.RefreshToken,
		Expires:  time.Now().Add(24 * 7 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	return c.NoContent(http.StatusOK)
}
