package main

import (
	"github.com/gofiber/fiber/v2"
	"gopi/config"
	"gopi/handlers"
	"log"
)

func main() {
	config.ReadConfig()

	app := fiber.New()

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
