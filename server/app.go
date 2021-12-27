package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rodny/image-saver/photo"
	photohttp "rodny/image-saver/photo/delivery/http"
	photousecase "rodny/image-saver/photo/usecase"
	"time"
)

type App struct {
	httpServer *http.Server

	photoUseCase photo.UseCase
}

func NewApp() *App {
	return &App{
		photoUseCase: photousecase.NewPhotoUseCase(),
	}
}

func (a *App) Run(port string) error {
	// Init gin handler
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
		CORSMiddleware(),
	)

	api := router.Group("/api")
	photohttp.RegisterHTTPEndpoints(api, a.photoUseCase)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
