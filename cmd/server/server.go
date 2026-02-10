package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func StartServer() {
	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	if err := e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
