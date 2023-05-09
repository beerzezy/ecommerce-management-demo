package common

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/beerzezy/ecommerce-management-demo/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func StartServer(e *echo.Echo, cfg config.Config) {
	port := fmt.Sprintf(":%d", cfg.Server.Port)
	if err := e.Start(port); err != nil {
		log.Info("shutting down the server")
	}
}

func WaitForGracefulShutdown(e *echo.Echo) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
