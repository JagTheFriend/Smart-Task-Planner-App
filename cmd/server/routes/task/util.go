package task

import (
	"smart-task-planner/cmd/middleware"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
)

func GetJwtDataMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		user, err := echo.ContextGet[*jwt.Token](c, "user")
		if err != nil {
			return echo.ErrUnauthorized.Wrap(err)
		}
		claims := user.Claims.(*middleware.JwtCustomClaims)
		c.Set("UserId", claims.UserId)
		return next(c)
	}
}
