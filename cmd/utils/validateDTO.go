package utils

import (
	"errors"
	"reflect"

	"github.com/labstack/echo/v5"
)

func BindAndValidate(c *echo.Context, dto any) error {
	// Must be a pointer to struct
	v := reflect.ValueOf(dto)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return errors.New("dto must be a non-nil pointer")
	}

	// Bind request body
	if err := c.Bind(dto); err != nil {
		return err
	}

	// Validate struct
	if err := c.Validate(dto); err != nil {
		return err
	}

	return nil
}
