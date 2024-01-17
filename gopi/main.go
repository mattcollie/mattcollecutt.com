package main

import (
	"fmt"
	"github.com/a-h/templ"
	"net/http"
)

func main() {
	component := page("Matt Collecutt")

	http.Handle("/", templ.Handler(component))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
