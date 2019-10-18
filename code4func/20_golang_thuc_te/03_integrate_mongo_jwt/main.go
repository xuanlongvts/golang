package main

import (
	"fmt"
	"net/http"

	//models "go-jwt/model"
	//repoImpl "go-jwt/repository/repoimpl"
	//"go-jwt/config"

	"go-jwt/handler"
	"go-jwt/driver"
)

func main() {
	fmt.Println("JWT")
	//mongo := driver.ConnectMongoDB()

	//userRepo := repoImpl.NewUserRepo(mongo.Client.Database(config.DB_NAME))
	//user := models.User{
	//	Email: "nongno@gmail.com",
	//	Password: "123456",
	//	DisplayName: "nongno",
	//}
	//err := userRepo.Insert(user)
	//if err == nil {
	//	fmt.Println("Insert ok")
	//}

	//user, _ := userRepo.FindUserByEmail("nongno@gmail.com")
	//fmt.Println("find user: ", user)

	//userLogin, _ := userRepo.CheckLoginInfo("nongno@gmail.com", "123456")
	//fmt.Println("Check login: ", userLogin)

	driver.ConnectMongoDB()
	http.HandleFunc("/register", handler.Register)

	fmt.Println("Server running [:8000]")
	http.ListenAndServe(":8000", nil)
}