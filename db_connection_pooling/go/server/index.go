package server

import (
	"db_cnn_pooling/db"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"time"
)

type Server struct {
	Router     *gin.Engine
	ConnNormal *sqlx.DB
	ConnPooled *sqlx.DB
}

var prefixDb = "DATABASE.postgres"
var driverName = "postgres"

func InitServer() Server {
	dsn := db.NewPostgreSql()

	maxConn := viper.GetInt(prefixDb + ".max_connections")
	idleConn := viper.GetInt(prefixDb + ".max_connections")
	lifeTimeConn := viper.GetInt(prefixDb + ".max_conn_lifetime")

	// ---- Connection Normal
	connNormal, err := sqlx.Open(driverName, dsn)
	if err != nil {
		log.Fatal("Connect to postgres failed: Normal: ", err.Error())
	}
	connNormal.SetMaxIdleConns(1)

	// ---- Connection pooled
	connPooled, err := sqlx.Open(driverName, dsn)
	if err != nil {
		log.Fatal("Connect to postgres failed: Pooled: ", err.Error())
	}
	connPooled.SetMaxOpenConns(maxConn)
	connPooled.SetMaxIdleConns(idleConn)
	connPooled.SetConnMaxLifetime(time.Duration(lifeTimeConn) * time.Minute)

	return Server{
		Router:     gin.Default(),
		ConnNormal: connNormal,
		ConnPooled: connPooled,
	}
}
