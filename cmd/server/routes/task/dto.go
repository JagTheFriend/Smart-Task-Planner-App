package task

import "time"

type CreateTaskDTO struct {
	Title       string    `json:"title" validate:"required,min=3,max=100"`
	Description string    `json:"description" validate:"omitempty,max=500"`
	Deadline    time.Time `json:"deadline" validate:"required"`
}

type UpdateTaskDTO struct {
	ID          uint      `json:"id" validate:"required,gt=0"`
	Title       string    `json:"title" validate:"omitempty,min=3,max=100"`
	Description string    `json:"description" validate:"omitempty,max=500"`
	Deadline    time.Time `json:"deadline" validate:"omitempty"`
	Completed   *bool     `json:"completed" validate:"omitempty"`
}

type DeleteTaskDTO struct {
	ID uint `query:"id" validate:"required,gt=0"`
}
