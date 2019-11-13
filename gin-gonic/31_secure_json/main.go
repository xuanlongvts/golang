package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// You can also use your own secure json prefix
	// r.SecureJsonPrefix(")]}',\n")

	r.GET("/somejson", func(c *gin.Context) {
		names := []string{"lerna", "austin", "foo"}

		// Will output  :   while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
	})

	r.Run()
}