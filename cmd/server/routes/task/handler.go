package task

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type TaskHandler struct {
	e  *echo.Group
	db *gorm.DB
}

func NewTaskHandler(e *echo.Group, db *gorm.DB) *TaskHandler {
	authGroup := e.Group("/task")

	authGroup.Use(GetJwtDataMiddleware)

	return &TaskHandler{e: authGroup, db: db}
}

func (h *TaskHandler) RegisterRoutes() {
	h.e.POST("", h.createTask)
	h.e.GET("", h.getTasks)
	h.e.PUT("", h.updateTask)
	h.e.DELETE("", h.deleteTask)
}

func (h *TaskHandler) createTask(c *echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("%d", c.Get("UserId")))
}

func (h *TaskHandler) getTasks(c *echo.Context) error {
	return nil
}

func (h *TaskHandler) updateTask(c *echo.Context) error {
	return nil
}

func (h *TaskHandler) deleteTask(c *echo.Context) error {
	return nil
}
