package router

import (
	"backend_github_trending/handler"
	"backend_github_trending/middleware"
	"github.com/labstack/echo"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API) SetupRouter() {
	api.Echo.POST("/user/sing-up", api.UserHandler.HandleSignUp)
	api.Echo.POST("/user/sing-in", api.UserHandler.HandleSignIn)
	api.Echo.GET("/user/profile", api.UserHandler.Profile, middleware.JwtMiddleware())
}
