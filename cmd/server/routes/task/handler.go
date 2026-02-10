package task

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type TaskHandler struct {
	e  *echo.Group
	db *gorm.DB
}

func NewTaskHandler(e *echo.Group, db *gorm.DB) *TaskHandler {
	authGroup := e.Group("/task")

	return &TaskHandler{e: authGroup, db: db}
}

func (h *TaskHandler) RegisterRoutes() {
}
