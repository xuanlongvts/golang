package router

import (
	"github.com/labstack/echo"
	"backend_github_trending/handler"
)

type API struct {
	Echo *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API) SetupRouter() {
	api.Echo.POST("/user/sing-up", api.UserHandler.HandleSignUp)
}