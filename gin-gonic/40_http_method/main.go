package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "GET")
	})
	r.POST("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "POST")
	})
	r.PUT("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "PUT")
	})
	r.DELETE("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "DELETE")
	})
	r.PATCH("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "PATCH")
	})
	r.HEAD("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "HEAD")
	})
	r.OPTIONS("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "OPTIONS")
	})

	r.Run()
}