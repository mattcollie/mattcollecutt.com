package main

import (
	"encoding/json"
	"fmt"
	"github.com/a-h/templ"
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

	socialLinks := []SocialLink{
		{name: "Github", url: "https://github.com/mattcollie", svg: "/static/images/github.svg"},
		{name: "Linkedin", url: "https://www.linkedin.com/in/matt-collecutt", svg: "/static/images/linkedin.svg"},
		{name: "Instagram", url: "https://www.instagram.com/matthewcollecutt", svg: "/static/images/instagram.svg"},
	}
	homeComponent := HomePage(socialLinks)
	http.Handle("/", templ.Handler(homeComponent))
	http.Handle("/photos", templ.Handler(PhotoPage(socialLinks)))
	http.Handle("/writing", templ.Handler(page("Writing", socialLinks)))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.Handle("/robots.txt", http.FileServer(http.Dir("./static/")))

	mediaResponse, err := media()
	if err != nil {
		log.Fatalf("Failed to get media: %v", err)
	}
	http.Handle("/instagram/media", templ.Handler(Photos(mediaResponse)))

	addr := cfg.Server.Host + ":" + cfg.Server.Port
	fmt.Println("Listening on " + addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Fatal error failed to start server: %s \\n", err)
	}
}

func media() ([]MediaItem, error) {
	url := fmt.Sprintf("https://graph.instagram.com/%s/media?fields=caption,id,media_type,media_url,timestamp&access_token=%s", cfg.Instagram.UserId, cfg.Instagram.AccessToken)

	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code %d", response.StatusCode)
	}
	defer response.Body.Close()

	var myResponse MediaResponse
	if err = json.NewDecoder(response.Body).Decode(&myResponse); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return myResponse.MediaItems, nil
}
