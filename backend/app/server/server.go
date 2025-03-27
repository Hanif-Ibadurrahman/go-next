package server

import (
	"backend/app/config"
	customMW "backend/app/middleware/logging"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitEcho() *echo.Echo {
	e := echo.New()
	e.Use(
		middleware.Recover(),
		middleware.CORS(),
		//middleware.Logger(),
		customMW.Logging(),
	)

	e.GET("/", func(c echo.Context) error {
		return ResponseStatusOK(c, "API service is running...", nil, nil, nil)
	})

	return e
}

// Start server
func Start(e *echo.Echo) {
	var (
		addr              = fmt.Sprintf(":%v", config.GetConfig().HTTPPort)
		HTTPServerTimeout = config.GetConfig().HTTPServerTimeOut
		readTime          = 1 * HTTPServerTimeout
		writeTime         = 20 * HTTPServerTimeout
	)

	s := &http.Server{
		Addr:         addr,
		ReadTimeout:  time.Second * time.Duration(readTime),
		WriteTimeout: time.Second * time.Duration(writeTime),
	}

	// Start server
	go func() {
		if err := e.StartServer(s); err != nil {
			e.Logger.Info("Shutting down the server ", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
