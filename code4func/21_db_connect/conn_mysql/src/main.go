package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go connect mysql")
	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/dbTest")
	checkErr(err)

	stmt, err := db.Prepare("insert users set user=?, email=?")
	checkErr(err)

	res, err := stmt.Exec("Lê Quang Định", "lqd@gmail.com")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("id inserted = ", id)

	// update
	stmt, err = db.Prepare("update users set user=? where id=?")
	checkErr(err)

	res, err = stmt.Exec("Long Lê Xuân", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("affect ", affect)
	fmt.Println("================")

	// query
	rows, err := db.Query("select * from users")
	checkErr(err)

	for rows.Next() {
		var id int
		var user string
		var email string
		err = rows.Scan(&id, &user, &email)
		checkErr(err)
		fmt.Println("id = ", id)
		fmt.Println("user = ", user)
		fmt.Println("email = ", email)
	}

	// delete
	stmt, err = db.Prepare("delete from users where id=?")
	checkErr(err)

	res, err = stmt.Exec(6)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Println("affect ", affect)

	db.Close()
}

func checkErr(err error)  {
	if err != nil {
		panic(err)
	}
}

// go get github.com/go-sql-driver/mysql     for download package github.com/go-sql-driver/mysql