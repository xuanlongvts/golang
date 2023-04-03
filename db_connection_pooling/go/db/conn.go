package db

import (
	"fmt"
	"github.com/spf13/viper"
)

func NewPostgreSql() string {
	var dbPrefix = "DATABASE.postgres."
	dbUser := viper.GetString(dbPrefix + "user")
	dbPassword := viper.GetString(dbPrefix + "pass")
	dbName := viper.GetString(dbPrefix + "name")
	dbHost := viper.GetString(dbPrefix + "host")
	dbPort := viper.GetInt(dbPrefix + "port")

	//var dsn = "postgres://postgres:password1@localhost:5432/postgres?sslmode=disable"
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	return dsn
}
