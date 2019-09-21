package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go connect mysql")
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/ginDB")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	insert, err := db.Query("INSERT INTO staff VALUES('4', 'Nguyen Van Minh', 'Long An')")
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	fmt.Println("Insert success")
}

// go get github.com/go-sql-driver/mysql     for download package github.com/go-sql-driver/mysql