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
	// user
	api.Echo.POST("/user/sing-up", api.UserHandler.HandleSignUp)
	api.Echo.POST("/user/sing-in", api.UserHandler.HandleSignIn)

	// profile
	user := api.Echo.Group("/user",  middleware.JwtMiddleware())
	user.GET("/profile", api.UserHandler.Profile)
	user.PUT("/update", api.UserHandler.UpdateProfile)
}
