package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// gin.H is a shortcut for map[string]interface{}
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": "hey",
		})
	})

	r.GET("/morejson", func(c *gin.Context) {
		var msg struct{
			Name string `json:"user"`
			Message string
			Number int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// Note that msg.Name becomes "user" in the JSON
		// Will output  :   {"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": "hey",
		})
	})

	r.GET("/yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": "hey",
		})
	})

	//r.GET("protobuf", func(c *gin.Context) {
	//	resp := []int64{int64(1), int64(2)}
	//	label := "test"
	//
	//	data := &protoexample.Test{
	//		Label: &label,
	//		Reps:  resp,
	//	}
	//
	//	c.ProtoBuf(http.StatusOK, data)
	//})

	r.Run()
}