package middleware

import (
	"backend_github_trending/log"
	"backend_github_trending/model"
	"backend_github_trending/model/req"
	"github.com/labstack/echo"
	"net/http"
)

func IsAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := req.ReqSignin{}
			if err := c.Bind(&req); err != nil {
				log.Error(err.Error())
				return c.JSON(http.StatusBadRequest, model.Response{
					StatusCode: http.StatusBadRequest,
					Message:    err.Error(),
					Data:       nil,
				})
			}

			if req.Email != "admin@gmail.com" {
				return c.JSON(http.StatusBadRequest, model.Response{
					StatusCode: http.StatusBadRequest,
					Message:    "You are unauthorized call this api",
					Data:       nil,
				})
			}

			return next(c)
		}
	}
}
