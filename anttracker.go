package main

import (
	"context"
	"net"
	"net/http"
	"time"
	"os"
	"os/signal"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	// Create our Echo instance.
	e := echo.New()

	// Setup our custom logger variables and configuration.
	e.Logger.SetLevel(log.INFO)

	if l, ok := e.Logger.(*log.Logger); ok {
		l.SetHeader("[${time_rfc3339}] [${level}]")
	  }

	// Let the user know the server is starting and setup.
	e.Logger.Info("Starting " + getenv("NAME", "antTracker") + "...")

	// Setup custom server based on our env values.
	antrackerServer := &http.Server {
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	l, _ := net.Listen(
		getenv("PROTOCOL", "tcp"),
		getenv("HOST", "") + ":" + getenv("PORT", "1337"),
	)
	e.Listener = l
	e.HideBanner = true

	// Setup any middleware needed.
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/", handleAnnounce)

	// We create a function for graceful shutdown if for some reason
	// our antTracker is in the middle of a transaction or request.
	go func() {
		if err := e.StartServer(antrackerServer); err != nil {
			e.Logger.Info("Gracefully shutting down " + getenv("NAME", "antTracker"))
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
