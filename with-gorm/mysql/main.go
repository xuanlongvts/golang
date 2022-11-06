package main

import (
	"mod-go-gorm/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	userRepo := controller.New()
	r.POST("/users", userRepo.CreateUser)
	r.GET("/users", userRepo.GetUsers)
	r.GET("/users/:email", userRepo.GetUser)
	r.PUT("/users/:email", userRepo.UpdateUser)
	r.DELETE("/users/:email", userRepo.DeleteUser)

	return r
}

func main() {
	r := setupRouter()

	_ = r.Run(":8000")
}
