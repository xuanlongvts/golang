package main

import "github.com/gin-gonic/gin"

func main() {
	//gin.DisableConsoleColor()
	gin.ForceConsoleColor()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context){
		c.String(200, "pong")
	})
	r.Run()
}