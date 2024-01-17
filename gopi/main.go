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
	component := page("Matt Collecutt")
	http.Handle("/", templ.Handler(component))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	socialLinks := []SocialLink{
		{name: "Github", url: "https://github.com/mattcollie", svg: "/static/images/github.svg"},
		{name: "Linkedin", url: "https://www.linkedin.com/in/matt-collecutt", svg: "/static/images/linkedin.svg"},
		{name: "Instagram", url: "https://www.instagram.com/matthewcollecutt", svg: "/static/images/instagram.svg"},
	}
	socialsComponent := socials(socialLinks)
	http.Handle("/socials", templ.Handler(socialsComponent))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
