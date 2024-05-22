package application

import (
	"context"
	"net/http"
	"os"

	"filestorage/internal/lib/logger"

	"github.com/gin-gonic/gin"
)

type App struct {
	log    logger.Logger
	server *http.Server
}

func NewApplication(log logger.Logger) *App {
	handler := gin.Default()

	addr := ":" + os.Getenv("PORT")
	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}
	return &App{
		log:    log,
		server: server,
	}
}

func (a *App) Start() {
	a.log.Info("starting server")
	a.log.Fatal(a.server.ListenAndServe())
}

func (a *App) Stop() {
	a.log.Info("stopping server")
	a.log.Fatal(a.server.Shutdown(context.Background()))
}
