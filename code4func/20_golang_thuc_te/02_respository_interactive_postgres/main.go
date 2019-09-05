package main

import (
	"fmt"
	"go-postgres/driver"
	"go-postgres/repository/repoImpl"
	models "go-postgres/model"
)

// host, port, user, password, dbname
const (
	host		= "localhost"
	port		= "5432"
	user		= "postgres"
	password 	= "123456"
	dbname		= "code4fun"
)

func main() {
	db := driver.Connect(host, port, user, password, dbname)

	err := db.SQL.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connect Success")

	userRepo := repoImpl.NewUserRepo(db.SQL)

	uhp := models.User {
		ID: 4,
		Name: "Ung Hoang Phuc",
		Gender: "Male",
		Email: "uhp@gmail.com",
	}

	myTam := models.User {
		ID: 5,
		Name: "My Tam",
		Gender: "Female",
		Email: "mt@gmail.com",
	}

	userRepo.Insert(uhp)
	userRepo.Insert(myTam)

	users, _ := userRepo.Select()

	for i := range users {
		fmt.Println(users[i])
	}
}

// Run command line
/*
	1. go mod init go-postgres
	2. go build
*/