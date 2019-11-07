package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router :=  gin.Default()

	// simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/login", func(c *gin.Context) {
			c.String(200, "Login")
		})
		v1.GET("/submit", func(c *gin.Context) {
			c.String(200, "Submit")
		})
		v1.GET("/read", func(c *gin.Context) {
			c.String(200, "Read")
		})
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", func(c *gin.Context) {
			c.String(200, "Login")
		})
		v2.GET("/submit", func(c *gin.Context) {
			c.String(200, "Submit")
		})
		v2.POST("/read", func(c *gin.Context) {
			c.String(200, "Read")
		})
	}

	router.Run()
}