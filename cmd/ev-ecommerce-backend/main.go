package main

import (
	"context"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/config"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	if err := godotenv.Load(); err != nil {
		println("error loading .env file")
	}

	cfg := config.NewConfig()
	e := app.Start(cfg)

	go func() {
		if err := e.Start(":" + cfg.Address); err != nil && err != http.ErrServerClosed {
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	println("💥 shutdown server ...")
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}


}
