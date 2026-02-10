package middleware

import (
	"fmt"
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.Validator.Struct(i); err != nil {
		slog.Error(fmt.Sprintf("Validator: %s", err.Error()))
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}
