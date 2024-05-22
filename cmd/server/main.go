package main

import (
	"filestorage/internal/application"
	"filestorage/internal/lib/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log := logger.NewLogger()
	app := application.NewApplication(log)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go app.Start()

	<-signalChan

	app.Stop()
}
