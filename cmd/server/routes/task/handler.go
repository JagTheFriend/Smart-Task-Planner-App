package task

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"smart-task-planner/cmd/utils"
	"smart-task-planner/internal/models"

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
	h.e.POST("", h.createTask)
	h.e.GET("", h.getTasks)
	h.e.PUT("", h.updateTask)
	h.e.DELETE("", h.deleteTask)
}

func (h *TaskHandler) createTask(c *echo.Context) error {
	var dto CreateTaskDTO
	if err := utils.BindAndValidate(c, &dto); err != nil {
		return err
	}

	userID := c.Get("UserId").(uint)

	task := models.Task{
		Title:       dto.Title,
		Description: dto.Description,
		Deadline:    dto.Deadline,
		Completed:   false,
		UserID:      userID,
	}

	res := h.db.Create(&task)
	if res.Error != nil {
		slog.Error(fmt.Sprintf("Task | Create Error | %s", res.Error.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to create task",
		})
	}

	slog.Info("Task: New Task Created")
	return c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) getTasks(c *echo.Context) error {
	userID := c.Get("UserId").(uint)

	var tasks []models.Task
	res := h.db.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&tasks)

	if res.Error != nil {
		slog.Error(fmt.Sprintf("Task | Fetch Error | %s", res.Error.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to fetch tasks",
		})
	}

	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) updateTask(c *echo.Context) error {
	var dto UpdateTaskDTO
	if err := utils.BindAndValidate(c, &dto); err != nil {
		return err
	}

	userID := c.Get("UserId").(uint)

	var existingTask models.Task
	res := h.db.First(&existingTask, "id = ? AND user_id = ?", dto.ID, userID)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Task not found",
		})
	}

	if res.Error != nil {
		slog.Error(fmt.Sprintf("Task | Fetch For Update Error | %s", res.Error.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Something went wrong",
		})
	}

	// Partial update
	if dto.Title != "" {
		existingTask.Title = dto.Title
	}
	if dto.Description != "" {
		existingTask.Description = dto.Description
	}
	if !dto.Deadline.IsZero() {
		existingTask.Deadline = dto.Deadline
	}
	if dto.Completed != nil {
		existingTask.Completed = *dto.Completed
	}

	res = h.db.Save(&existingTask)
	if res.Error != nil {
		slog.Error(fmt.Sprintf("Task | Update Error | %s", res.Error.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to update task",
		})
	}

	slog.Info("Task: Task Updated")
	return c.JSON(http.StatusOK, existingTask)
}

func (h *TaskHandler) deleteTask(c *echo.Context) error {
	var dto DeleteTaskDTO
	if err := utils.BindAndValidate(c, &dto); err != nil {
		return err
	}

	userID := c.Get("UserId").(uint)

	res := h.db.
		Where("id = ? AND user_id = ?", dto.ID, userID).
		Delete(&models.Task{})

	if res.Error != nil {
		slog.Error(fmt.Sprintf("Task | Delete Error | %s", res.Error.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to delete task",
		})
	}

	if res.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Task not found",
		})
	}

	slog.Info("Task: Task Deleted")
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Task deleted successfully",
	})
}
