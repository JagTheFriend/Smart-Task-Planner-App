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
