package authentication

import (
	"errors"
	"log/slog"
	"net/http"
	"smart-task-planner/internal/models"

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
	var u SignUpDTO
	if err := c.Bind(&u); err != nil {
		return err
	}
	if err := c.Validate(&u); err != nil {
		return err
	}

	var existingUser models.User
	res := h.db.First(&existingUser, models.User{Email: u.Email})
	if res.Error != nil && !(errors.Is(res.Error, gorm.ErrRecordNotFound)) {
		return c.JSON(http.StatusConflict, map[string]string{"message": "User Already Exists"})
	}

	res = h.db.Create(&models.User{
		Email:    u.Email,
		Password: u.Password,
		Name:     u.Name,
	})

	if res.Error != nil {
		slog.Warn("Auth: Failed to Resgister User")
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to register"})
	}
	slog.Info("Auth:New User Detail Saved")
	return c.JSON(http.StatusCreated, map[string]string{"message": "User Registered"})
}

func (h *AuthHandler) login(c *echo.Context) error {
	var u LoginDTO
	if err := c.Bind(&u); err != nil {
		return err
	}
	if err := c.Validate(&u); err != nil {
		return err
	}

	var existingUser models.User
	res := h.db.First(&existingUser, models.User{Email: u.Email})
	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusConflict, map[string]string{"message": "User Does Not Exist"})
	}

	return c.JSON(http.StatusOK, map[string]models.User{"message": existingUser})
}
