package handler

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mattcollie/mattcollecutt.com/config"
	"github.com/mattcollie/mattcollecutt.com/model"
	"github.com/mattcollie/mattcollecutt.com/view/photo"
	"net/http"
)

type PhotoHandler struct{}

func (h PhotoHandler) HandlePhotoShow(c echo.Context) error {
	rsp, err := media()
	if err != nil {
		return err
	}
	return render(c, photo.Show(rsp))
}

func (h PhotoHandler) HandlePhotoMedia(c echo.Context) error {
	rsp, err := media()
	if err != nil {
		return err
	}
	return render(c, photo.Photos(rsp))
}

func (h PhotoHandler) HandlePhotoMediaAfter(c echo.Context) error {
	after := c.Param("after")
	rsp, err := mediaAfter(after)
	if err != nil {
		return err
	}
	return render(c, photo.Photos(rsp))
}

func media() (model.MediaResponse, error) {
	limit := 12
	url := fmt.Sprintf("https://graph.instagram.com/%s/media?fields=caption,id,media_type,media_url,timestamp,children{id,media_url}&access_token=%s&limit=%d", config.Cfg.Instagram.UserId, config.Cfg.Instagram.AccessToken, limit)

	response, err := http.Get(url)
	if err != nil {
		return model.MediaResponse{}, fmt.Errorf("error making request: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return model.MediaResponse{}, fmt.Errorf("API request failed with status code %d", response.StatusCode)
	}
	defer response.Body.Close()

	var myResponse model.MediaResponse
	if err = json.NewDecoder(response.Body).Decode(&myResponse); err != nil {
		return model.MediaResponse{}, fmt.Errorf("error decoding response: %w", err)
	}

	return myResponse, nil
}

func mediaAfter(after string) (model.MediaResponse, error) {
	limit := 12
	url := fmt.Sprintf("https://graph.instagram.com/%s/media?fields=caption,id,media_type,media_url,timestamp,children{id,media_url}&access_token=%s&limit=%d&after=%s", config.Cfg.Instagram.UserId, config.Cfg.Instagram.AccessToken, limit, after)

	response, err := http.Get(url)
	if err != nil {
		return model.MediaResponse{}, fmt.Errorf("error making request: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return model.MediaResponse{}, fmt.Errorf("API request failed with status code %d", response.StatusCode)
	}
	defer response.Body.Close()

	var myResponse model.MediaResponse
	if err = json.NewDecoder(response.Body).Decode(&myResponse); err != nil {
		return model.MediaResponse{}, fmt.Errorf("error decoding response: %w", err)
	}

	return myResponse, nil
}
