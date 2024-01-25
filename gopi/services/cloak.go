package services

import (
	"context"
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"log"

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

// GetLoginURL generates the Keycloak login URL
func GetLoginURL(c *fiber.Ctx) error {
	loginURL := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/auth?client_id=%s&redirect_uri=%s&response_type=code&scope=openid&kc_idp_hint=github", keycloakURL, realm, clientID, redirectURI)

	if c.Get("HX-Request") == "true" {
		c.Set("HX-Redirect", loginURL)
		return c.SendStatus(fiber.StatusOK)
	}

	return c.Redirect(loginURL, fiber.StatusTemporaryRedirect)
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

	return c.JSON(token)
}
