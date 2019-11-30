package server

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/cmelgarejo/go-gql-server/internal/handlers"
	"github.com/cmelgarejo/go-gql-server/pkg/utils"
)

var HOST, PORT string

func init() {
	HOST = utils.MustGet("GQL_SERVER_HOST")
	PORT = utils.MustGet("GQL_SERVER_PORT")
}

func Run() {
	r := gin.Default()
	r.GET("/ping", handlers.Ping())
	log.Println("Running @ http://" + HOST + ":" + PORT )
	log.Fatalln(r.Run(HOST + ":" + PORT))
}