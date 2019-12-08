package main

import (
	"backend_github_trending/db"
	"backend_github_trending/log"
	"context"
	"github.com/labstack/echo"
	"net/http"
	"os"
)

func init() {
	os.Setenv("APP_NAME", "github")
	log.InitLogger(false)
}

func main() {

	sql := &db.Sql{
		Host: "localhost",
		Port: 5432,
		UserName: "postgres",
		Password: "123456",
		DbName: "golang-flutter",
	}
	sql.Connect()
	defer sql.Close()

	var email string
	err := sql.Db.GetContext(context.Background(), &email, "select email from users where email=$1", "abc@gmail.com")
	if err != nil {
		log.Error(err.Error())
	}



	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
