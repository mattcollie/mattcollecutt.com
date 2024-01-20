package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

type SocialLink struct {
	name string
	url  string
	svg  string
}

func main() {
	socialLinks := []SocialLink{
		{name: "Github", url: "https://github.com/mattcollie", svg: "/static/images/github.svg"},
		{name: "Linkedin", url: "https://www.linkedin.com/in/matt-collecutt", svg: "/static/images/linkedin.svg"},
		{name: "Instagram", url: "https://www.instagram.com/matthewcollecutt", svg: "/static/images/instagram.svg"},
	}
	homeComponent := HomePage(socialLinks)
	http.Handle("/", templ.Handler(homeComponent))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	fmt.Println("Listening on :3000")
	http.ListenAndServe("127.0.0.1:3000", nil)
}
