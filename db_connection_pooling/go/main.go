package main

import (
	"db_cnn_pooling/config"
	"db_cnn_pooling/routers"
	"db_cnn_pooling/server"
	"embed"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var f embed.FS

func main() {
	config.LoadConfig()
	server := server.InitServer()

	server.Router.StaticFile("/public/favicon.ico", "./html/favicon.ico")
	server.Router.StaticFile("/", "./html/index.html")

	v1 := server.Router.Group("/products")
	v1.GET("/new", routers.GetProductNew())
	v1.GET("/normal", routers.GetProductNormal(server))
	v1.GET("/pooled", routers.GetProductPooled(server))

	// Start the HTTP server
	port := fmt.Sprintf(":%d", viper.GetInt("HTTP_SERVER.port"))
	if err := server.Router.Run(port); err != nil {
		log.Fatalf("Unable to start HTTP server: %v\n", err)
	}
}
