package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mattcollie/mattcollecutt.com/config"
	"github.com/mattcollie/mattcollecutt.com/handler"
	"github.com/mattcollie/mattcollecutt.com/middleware"
	"github.com/mattcollie/mattcollecutt.com/service"
	"log"
)

func main() {
	config.Read()
	service.StartCloak()

	app := echo.New()

	app.Use(middleware.WithAuthentication)

	authHandler := handler.AuthHandler{}
	app.GET("/login", authHandler.HandleAuthLogin)
	app.GET("/logout", authHandler.HandleAuthLogout)
	app.GET("/callback", authHandler.HandleAuthCallback)
	app.GET("/auth/refresh", authHandler.HandleAuthRefresh)

	rootHandler := handler.RootHandler{}
	app.GET("/", rootHandler.HandleRootShow)

	photoHandler := handler.PhotoHandler{}
	app.GET("/photos", photoHandler.HandlePhotoShow)
	app.GET("/media", photoHandler.HandlePhotoMedia)
	app.GET("/media/after/:after", photoHandler.HandlePhotoMediaAfter)

	app.Static("/static", "./static")

	addr := config.Cfg.Server.Host + ":" + config.Cfg.Server.Port
	log.Println("Listening on " + addr)
	if err := app.Start(addr); err != nil {
		log.Fatalf("Fatal error failed to start server: %s\n", err)
	}
}
