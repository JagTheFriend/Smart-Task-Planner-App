package server

import (
	"fmt"
	"net/http"
	"os"
	"smart-task-planner/cmd/server/routes/authentication"
	"smart-task-planner/internal/database"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func StartServer() {
	db := database.Connect()
	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	apiGroup := e.Group("/api/v1")
	apiGroup.GET("/health", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"health": "ok"})
	})

	authHandler := authentication.NewAuthHandler(apiGroup, db)
	authHandler.RegisterRoutes()

	if err := e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
