package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db, err := sql.Open("mysql", "root:12345678@tcp(localhost:3306)/dev_db_auth?multiStatements=true")
	if err != nil {

		panic("Cannot connect to mysql")
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic("mysql.WithInstance")
	}

	m, err := migrate.NewWithDatabaseInstance("file:///migrations", "mysql", driver)
	if err != nil {
		panic("migrate.NewWithDatabaseInstance")
	}
	fmt.Println("---> m: ", m)

	m.Steps(1)
}
