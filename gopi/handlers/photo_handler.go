package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gopi/config"
	"gopi/models"
	"gopi/templates"
	"net/http"
)

func PhotoHandler(c *fiber.Ctx) error {
	mediaResponse, err := media()
	if err != nil {
		return err
	}
	return Render(c, templates.PhotoPage(mediaResponse))
}

func MediaHandler(c *fiber.Ctx) error {
	mediaResponse, err := media()
	if err != nil {
		return err
	}
	return Render(c, templates.Photos(mediaResponse))
}

func MediaAfterHandler(c *fiber.Ctx) error {
	after := c.Params("after")
	mediaResponse, err := mediaAfter(after)
	if err != nil {
		return err
	}
	return Render(c, templates.Photos(mediaResponse))
}

func media() (models.MediaResponse, error) {
	limit := 12
	url := fmt.Sprintf("https://graph.instagram.com/%s/media?fields=caption,id,media_type,media_url,timestamp,children{id,media_url}&access_token=%s&limit=%d", config.Cfg.Instagram.UserId, config.Cfg.Instagram.AccessToken, limit)

	response, err := http.Get(url)
	if err != nil {
		return models.MediaResponse{}, fmt.Errorf("error making request: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return models.MediaResponse{}, fmt.Errorf("API request failed with status code %d", response.StatusCode)
	}
	defer response.Body.Close()

	var myResponse models.MediaResponse
	if err = json.NewDecoder(response.Body).Decode(&myResponse); err != nil {
		return models.MediaResponse{}, fmt.Errorf("error decoding response: %w", err)
	}

	return myResponse, nil
}

func mediaAfter(after string) (models.MediaResponse, error) {
	limit := 12
	url := fmt.Sprintf("https://graph.instagram.com/%s/media?fields=caption,id,media_type,media_url,timestamp,children{id,media_url}&access_token=%s&limit=%d&after=%s", config.Cfg.Instagram.UserId, config.Cfg.Instagram.AccessToken, limit, after)

	response, err := http.Get(url)
	if err != nil {
		return models.MediaResponse{}, fmt.Errorf("error making request: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return models.MediaResponse{}, fmt.Errorf("API request failed with status code %d", response.StatusCode)
	}
	defer response.Body.Close()

	var myResponse models.MediaResponse
	if err = json.NewDecoder(response.Body).Decode(&myResponse); err != nil {
		return models.MediaResponse{}, fmt.Errorf("error decoding response: %w", err)
	}

	return myResponse, nil
}
