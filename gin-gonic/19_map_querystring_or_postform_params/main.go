package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v", ids, names)
	})

	r.Run()
}

// Note: http://localhost:8080/post?ids[a]=1234&ids[b]=hello

// or post form names[first]=thinkerou&names[second]=tianou