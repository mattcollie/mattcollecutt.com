package service

import (
	"context"
	"github.com/Nerzal/gocloak/v13"
	"github.com/mattcollie/mattcollecutt.com/config"
	"log"
)

var client *gocloak.GoCloak

func StartCloak() {
	client = gocloak.NewClient(config.Cfg.Authentication.KeycloakURL)
}

func ExchangeToken(code string) (*gocloak.JWT, error) {
	ctx := context.Background()
	tokenOptions := gocloak.TokenOptions{
		ClientID:     &config.Cfg.Authentication.ClientID,
		ClientSecret: &config.Cfg.Authentication.ClientSecret,
		Code:         &code,
		RedirectURI:  &config.Cfg.Authentication.RedirectURI,
		GrantType:    gocloak.StringP("authorization_code"),
	}

	token, err := client.GetToken(ctx, config.Cfg.Authentication.Realm, tokenOptions)
	if err != nil {
		log.Printf("Error exchanging code for token: %v\n", err)
		return nil, err
	}
	return token, nil
}

func ValidateAccessToken(accessToken string) error {
	ctx := context.Background()
	_, _, err := client.DecodeAccessToken(ctx, accessToken, config.Cfg.Authentication.Realm)
	return err
}

func ValidateCode(ctx context.Context, code string) (*gocloak.JWT, error) {
	tokenOptions := gocloak.TokenOptions{
		ClientID:     &config.Cfg.Authentication.ClientID,
		ClientSecret: &config.Cfg.Authentication.ClientSecret,
		Code:         &code,
		RedirectURI:  &config.Cfg.Authentication.RedirectURI,
		GrantType:    gocloak.StringP("authorization_code"),
	}

	token, err := client.GetToken(ctx, config.Cfg.Authentication.Realm, tokenOptions)
	if err != nil {
		log.Printf("Error exchanging code for token: %v\n", err)
		return token, err
	}

	_, _, err = client.DecodeAccessToken(ctx, token.AccessToken, config.Cfg.Authentication.Realm)
	return token, err
}

func GetRefresh(refreshToken string) (*gocloak.JWT, error) {
	newToken, err := client.RefreshToken(context.Background(), refreshToken, config.Cfg.Authentication.ClientID, config.Cfg.Authentication.ClientSecret, config.Cfg.Authentication.Realm)
	if err != nil {
		return newToken, err
	}
	return newToken, nil
}
