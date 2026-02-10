package authentication

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"smart-task-planner/cmd/middleware"
	"smart-task-planner/internal/models"

	"github.com/golang-jwt/jwt/v5"
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
	if !(errors.Is(res.Error, gorm.ErrRecordNotFound)) {
		return c.JSON(http.StatusConflict, map[string]string{"message": "User Already Exists"})
	}

	if res.Error != nil {
		slog.Error(fmt.Sprintf("Auth | Login Error | %s", res.Error.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Something went wrong"})
	}

	res = h.db.Create(&models.User{
		Email:    u.Email,
		Password: u.Password,
		Name:     u.Name,
	})

	if res.Error != nil {
		slog.Error(fmt.Sprintf("Auth | Signup Error | %s", res.Error.Error()))
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
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusConflict, map[string]string{"message": "User Does Not Exist"})
	}

	if res.Error != nil {
		slog.Error(fmt.Sprintf("Auth | Login Error | %s", res.Error.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Something went wrong"})
	}

	claims := middleware.JwtCustomClaims{
		UserId: existingUser.ID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"message": t})
}
