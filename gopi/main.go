package main

import (
	"encoding/json"
	"fmt"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"log"
	"net/http"
)

type SocialLink struct {
	name string
	url  string
	svg  string
}

func main() {
	readConfig()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		socialLinks := []SocialLink{
			{name: "Github", url: "https://github.com/mattcollie", svg: "/static/images/github.svg"},
			{name: "Linkedin", url: "https://www.linkedin.com/in/matt-collecutt", svg: "/static/images/linkedin.svg"},
			{name: "Instagram", url: "https://www.instagram.com/matthewcollecutt", svg: "/static/images/instagram.svg"},
		}
		return Render(c, HomePage(socialLinks))
	})
	app.Get("/media", func(c *fiber.Ctx) error {
		mediaResponse, err := media()
		if err != nil {
			return err
		}
		return Render(c, Photos(mediaResponse))
	})
	app.Get("/media/after/:after", func(c *fiber.Ctx) error {
		after := c.Params("after")
		mediaResponse, err := mediaAfter(after)
		if err != nil {
			return err
		}
		return Render(c, Photos(mediaResponse))
	})
	app.Get("/photos", func(c *fiber.Ctx) error {
		socialLinks := []SocialLink{
			{name: "Github", url: "https://github.com/mattcollie", svg: "/static/images/github.svg"},
			{name: "Linkedin", url: "https://www.linkedin.com/in/matt-collecutt", svg: "/static/images/linkedin.svg"},
			{name: "Instagram", url: "https://www.instagram.com/matthewcollecutt", svg: "/static/images/instagram.svg"},
		}
		mediaResponse, err := media()
		if err != nil {
			return err
		}
		return Render(c, PhotoPage(socialLinks, mediaResponse))
	})
	app.Get("/writing", func(c *fiber.Ctx) error {
		socialLinks := []SocialLink{
			{name: "Github", url: "https://github.com/mattcollie", svg: "/static/images/github.svg"},
			{name: "Linkedin", url: "https://www.linkedin.com/in/matt-collecutt", svg: "/static/images/linkedin.svg"},
			{name: "Instagram", url: "https://www.instagram.com/matthewcollecutt", svg: "/static/images/instagram.svg"},
		}
		return Render(c, page("Writing", socialLinks))
	})
	app.Get("/robots.txt", func(c *fiber.Ctx) error {
		return c.SendFile("./static/robots.txt")
	})
	app.Static("/static", "./static")

	addr := cfg.Server.Host + ":" + cfg.Server.Port
	log.Println("Listening on " + addr)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("Fatal error failed to start server: %s\n", err)
	}
}

func media() (MediaResponse, error) {
	limit := 12
	url := fmt.Sprintf("https://graph.instagram.com/%s/media?fields=caption,id,media_type,media_url,timestamp,children{id,media_url}&access_token=%s&limit=%d", cfg.Instagram.UserId, cfg.Instagram.AccessToken, limit)

	response, err := http.Get(url)
	if err != nil {
		return MediaResponse{}, fmt.Errorf("error making request: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return MediaResponse{}, fmt.Errorf("API request failed with status code %d", response.StatusCode)
	}
	defer response.Body.Close()

	var myResponse MediaResponse
	if err = json.NewDecoder(response.Body).Decode(&myResponse); err != nil {
		return MediaResponse{}, fmt.Errorf("error decoding response: %w", err)
	}

	return myResponse, nil
}

func mediaAfter(after string) (MediaResponse, error) {
	limit := 12
	url := fmt.Sprintf("https://graph.instagram.com/%s/media?fields=caption,id,media_type,media_url,timestamp,children{id,media_url}&access_token=%s&limit=%d&after=%s", cfg.Instagram.UserId, cfg.Instagram.AccessToken, limit, after)

	response, err := http.Get(url)
	if err != nil {
		return MediaResponse{}, fmt.Errorf("error making request: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return MediaResponse{}, fmt.Errorf("API request failed with status code %d", response.StatusCode)
	}
	defer response.Body.Close()

	var myResponse MediaResponse
	if err = json.NewDecoder(response.Body).Decode(&myResponse); err != nil {
		return MediaResponse{}, fmt.Errorf("error decoding response: %w", err)
	}

	return myResponse, nil
}

func Render(c *fiber.Ctx, component templ.Component, options ...func(handler *templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}

	return adaptor.HTTPHandler(componentHandler)(c)
}
