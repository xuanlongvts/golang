package main

import (
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		// you can bind multipart form with explicit binding declaration:
		// c.ShouldBindWith(&form, binding.Form)
		// or you can simply use autobinding with ShouldBind method:
		var form LoginForm
		if c.ShouldBind(&form) == nil {
			if form.User == "longle" && form.Password == "123" {
				c.JSON(200, gin.H{
					"status": "You are login in",
				})
			} else {
				c.JSON(401, gin.H{
					"status": "Unauthorized",
				})
			}
		}
	})

	router.Run()
}