package services

import (
	"context"
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"gopi/config"
)

var (
	keycloakURL  string
	realm        string
	clientID     string
	clientSecret string
	redirectURI  string
)

var client *gocloak.GoCloak

func Init() {
	keycloakURL = config.Cfg.Authentication.KeycloakURL
	realm = config.Cfg.Authentication.Realm
	clientID = config.Cfg.Authentication.ClientID
	clientSecret = config.Cfg.Authentication.ClientSecret
	redirectURI = config.Cfg.Authentication.RedirectURI

	client = gocloak.NewClient(keycloakURL)
}

func GetLoginURL(c *fiber.Ctx) error {
	loginURL := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/auth?client_id=%s&redirect_uri=%s&response_type=code&scope=openid&kc_idp_hint=github", keycloakURL, realm, clientID, redirectURI)

	if c.Get("HX-Request") == "true" {
		c.Set("HX-Redirect", loginURL)
		return c.SendStatus(fiber.StatusOK)
	}

	return c.Redirect(loginURL, fiber.StatusTemporaryRedirect)
}

func InvalidateJWTCookies(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_jwt",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})

	return c.SendString(`
		<button
			hx-get="/login"
			hx-boost="true"
			class="bg-neutral-750 rounded h-[30px] w-full text-xs font-semibold text-neutral-300 hover:text-neutral-50 active:bg-neutral-800">
			Sign in
		</button>`)
}

func HandleLoginCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	tokenOptions := gocloak.TokenOptions{
		ClientID:     &clientID,
		ClientSecret: &clientSecret,
		Code:         &code,
		RedirectURI:  &redirectURI,
		GrantType:    gocloak.StringP("authorization_code"),
	}

	token, err := client.GetToken(context.Background(), realm, tokenOptions)
	if err != nil {
		log.Printf("Error exchanging code for token: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error exchanging code for token: " + err.Error())
	}

	err = ValidateAccessToken(token.AccessToken)
	if err != nil {
		log.Printf("Error decoding access token: %v\n", err)
		return c.Status(fiber.StatusUnauthorized).SendString("Error decoding access token: " + err.Error())
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token.AccessToken,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
		Expires:  time.Now().Add(24 * time.Hour),
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_jwt",
		Value:    token.RefreshToken,
		Expires:  time.Now().Add(24 * 7 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})

	return c.Redirect("/")
}

func RefreshToken(c *fiber.Ctx) error {
	refreshTokenCookie := c.Cookies("refresh_jwt")

	newToken, err := client.RefreshToken(context.Background(), refreshTokenCookie, clientID, clientSecret, realm)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Failed to refresh token")
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    newToken.AccessToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_jwt",
		Value:    newToken.RefreshToken,
		Expires:  time.Now().Add(24 * 7 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})

	return c.SendStatus(fiber.StatusOK)
}

func ValidateAccessToken(accessToken string) error {
	ctx := context.Background()
	_, _, err := client.DecodeAccessToken(ctx, accessToken, realm)
	return err
}
