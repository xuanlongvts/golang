package main

import (
	"fmt"
	"go-jwt/driver"
	"go-jwt/config"
	models "go-jwt/model"
	repoImpl "go-jwt/repository/repoimpl"
)

func main() {
	fmt.Println("JWT")
	mongo := driver.ConnectMongoDB()

	userRepo := repoImpl.NewUserRepo(mongo.Client.Database(config.DB_NAME))
	user := models.User{
		Email: "nongno@gmail.com",
		Password: "123456",
		DisplayName: "nongno",
	}
	err := userRepo.Insert(user)
	if err == nil {
		fmt.Println("Insert ok")
	}
}