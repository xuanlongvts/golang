package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	//"time"
)

// users data
type users struct {
	id int
	name string
	gender string
	email string
}

const (
	DB_DRIVER_NAME = "postgres"
	DB_USER = "postgres"
	DB_PASSWORD = "123456"
	DB_NAME = "code4fun"
)

func main() {
	fmt.Println("Connect to postgres db")

	dbInfor := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open(DB_DRIVER_NAME, dbInfor)
	checkErr(err)
	defer db.Close()

	fmt.Println("# Insert values")
	//var lastInsertId int
	//err = db.QueryRow("INSERT INTO users(id, name, gender, email) VALUES($1, $2, $3, $4)", "9", "Lê Quang Định", "Male", "lqd@gmail.com").Scan(&lastInsertId)
	//checkErr(err)
	//fmt.Println("last insert id = ", lastInsertId)

	fmt.Println("# Updating")
	stmt, err := db.Prepare("update users set name=$1 where id=$2")
	checkErr(err)
	res, err := stmt.Exec("LQĐ update", 8)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect, "rows changed")

	fmt.Println("# Querying")
	rows, err := db.Query("SELECT * FROM users")
	checkErr(err)

	for rows.Next() {
		var id int
		var name string
		var gender string
		var email string
		err = rows.Scan(&id, &name, &gender, &email)
		checkErr(err)
		fmt.Println("id | name | gener | email ")
		fmt.Printf("%3v | %8v | %6v | %6v\n", id, name, gender, email)
	}

	fmt.Println("# Deleting")
	stmt, err = db.Prepare("delete from users where id=$1")
	checkErr(err)
	res, err = stmt.Exec(9)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect, "rows changed")
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}