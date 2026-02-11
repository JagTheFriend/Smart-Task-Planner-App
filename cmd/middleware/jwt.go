package middleware

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
)

type JwtCustomClaims struct {
	UserId uint `json:"userId"`
	jwt.RegisteredClaims
}

func JwtMiddleware() echo.MiddlewareFunc {
	// Configure middleware with the custom claims type
	config := echojwt.Config{
		NewClaimsFunc: func(c *echo.Context) jwt.Claims {
			return new(JwtCustomClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_KEY")),
	}

	return echojwt.WithConfig(config)
}

func GetJwtDataMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		user, err := echo.ContextGet[*jwt.Token](c, "user")
		if err != nil {
			return echo.ErrUnauthorized
		}

		claims, ok := user.Claims.(*JwtCustomClaims)
		if !ok {
			return echo.ErrUnauthorized
		}

		c.Set("UserId", claims.UserId)
		return next(c)
	}
}
