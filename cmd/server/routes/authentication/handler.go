package authentication

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type AuthHandler struct {
	e  *echo.Group
	db *gorm.DB
}

func NewAuthHandler(e *echo.Group, db *gorm.DB) *AuthHandler {
	authGroup := e.Group("/auth")

	return &AuthHandler{e: authGroup, db: db}
}

func (h *AuthHandler) RegisterRoutes() {
	h.e.POST("/signup", h.signup)
	h.e.POST("/login", h.login)
}

func (h *AuthHandler) signup(c *echo.Context) error {
	return nil
}

func (h *AuthHandler) login(c *echo.Context) error {
	return nil
}
