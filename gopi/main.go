package main

import (
	"fmt"
	"github.com/a-h/templ"
	"net/http"
)

func main() {
	component := hello("Matt")

	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
