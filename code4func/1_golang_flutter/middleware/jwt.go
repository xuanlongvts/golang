package middleware

import (
	"backend_github_trending/model"
	"backend_github_trending/security"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func JwtMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: []byte(security.SECRECT_KEY),
	}

	return middleware.JWTWithConfig(config)
}
