package main

import (
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
	http.Handle("/photos", templ.Handler(page("Photos", socialLinks)))
	http.Handle("/writing", templ.Handler(page("Writing", socialLinks)))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.Handle("/robots.txt", http.FileServer(http.Dir("./static/")))

	addr := cfg.Server.Host + ":" + cfg.Server.Port
	fmt.Println("Listening on " + addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Fatal error failed to start server: %s \\n", err)
	}
}
