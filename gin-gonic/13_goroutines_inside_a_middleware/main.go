package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/long_async", func(c *gin.Context) {
		cCp := c.Copy()

		go func() {
			time.Sleep(5 * time.Second)

			log.Println("Done! in path Asynchronous " + cCp.Request.URL.Path)
		}()
	})

	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)

		log.Println("Done! in path syn " + c.Request.URL.Path)
	})

	r.Run()
}

/*
When starting new Goroutines inside a middleware or handler, you SHOULD NOT use the original context inside it, you have to use a read-only copy.
*/