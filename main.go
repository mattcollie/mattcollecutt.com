package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/mattcollie/mattcollecutt.com/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Service struct{}

func main() {
	host := flag.String("host", "localhost", "Server host")
	port := flag.Int("port", 8080, "Server port")
	docker := flag.Bool("docker", false, "Run in docker")
	staticPath := flag.String("static", "./static", "Static path")
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	service := Service{}
	ginHtmlRenderer := router.HTMLRender
	router.HTMLRender = &handler.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}

	api := router.Group("/api")
	api.GET("/health", service.HealthCheck)
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	rootHandler := handler.RootHandler{}
	router.GET("/", rootHandler.HandleRootShow)

	router.Use(static.Serve("/static", static.LocalFile(*staticPath, true)))
	router.NoRoute(func(c *gin.Context) {
		c.File(fmt.Sprintf("%s/index.html", staticPath))
	})

	var serverPath string
	if *docker {
		serverPath = "0.0.0.0:8080"
		log.Println("Server started at http://0.0.0.0:8080 ...")
	} else {
		serverPath = fmt.Sprintf("%s:%d", *host, *port)
		log.Printf("Server started at http://%s ...\n", serverPath)
	}

	server := &http.Server{
		Addr:         serverPath,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}
	log.Println("Server exiting")
}

func (s *Service) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
