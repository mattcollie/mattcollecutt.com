package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mattcollie/mattcollecutt.com/handler"
	"log"
)

func main() {
	host := flag.String("host", "localhost", "Server host")
	port := flag.Int("port", 8080, "Server port")
	docker := flag.Bool("docker", false, "Run in docker")
	static := flag.String("static", "./static", "Static path")
	flag.Parse()

	app := echo.New()

	app.Static("/static", *static)

	rootHandler := handler.RootHandler{}
	app.GET("/", rootHandler.HandleRootShow)

	var addr string
	if *docker {
		addr = "0.0.0.0:8080"
		log.Println("Server started at http://0.0.0.0:8080 ...")
	} else {
		addr = fmt.Sprintf("%s:%d", *host, *port)
		log.Printf("Server started at http://%s ...\n", addr)
	}
	if err := app.Start(addr); err != nil {
		log.Fatalf("Fatal error failed to start server: %s\n", err)
	}
}
