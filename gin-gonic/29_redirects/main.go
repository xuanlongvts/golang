package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/testredirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://google.com/")
	})

	r.GET("/test", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"

		r.HandleContext(c)
	})

	r.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"Hello": "World"})
	})
	r.Run()
}