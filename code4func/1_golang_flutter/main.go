package main

import (
	"backend_github_trending/db"
	"backend_github_trending/handler"
	"backend_github_trending/log"
	"backend_github_trending/repository/repo_impl"
	"backend_github_trending/router"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

func init() {
	os.Setenv("APP_NAME", "github")
	log.InitLogger(false)
}

func main() {

	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "postgres",
		Password: "123456",
		DbName:   "golang-flutter",
	}
	sql.Connect()
	defer sql.Close()

	e := echo.New()
	e.Use(middleware.AddTrailingSlash())

	userHandler := handler.UserHandler{UserRepo: repo_impl.NewUserRepo(sql)}

	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
	}
	api.SetupRouter()

	e.Logger.Fatal(e.Start(":1323"))
}
