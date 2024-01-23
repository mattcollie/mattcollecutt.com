package main

import (
	"github.com/gofiber/fiber/v2"
	"gopi/config"
	"gopi/handlers"
	"gopi/services"
	"log"
)

func main() {
	config.ReadConfig()
	services.Init()

	app := fiber.New()

	// New routes for Keycloak integration
	app.Get("/login", services.GetLoginURL)
	app.Get("/callback", services.HandleLoginCallback)

	app.Get("/", handlers.RootHandler)
	app.Get("/media", handlers.MediaHandler)
	app.Get("/media/after/:after", handlers.MediaAfterHandler)
	app.Get("/photos", handlers.PhotoHandler)
	app.Get("/robots.txt", func(c *fiber.Ctx) error {
		return c.SendFile("./static/robots.txt")
	})
	app.Static("/static", "./static")
	app.Use(handlers.NotFoundHandler)

	addr := config.Cfg.Server.Host + ":" + config.Cfg.Server.Port
	log.Println("Listening on " + addr)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("Fatal error failed to start server: %s\n", err)
	}
}
