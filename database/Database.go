package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DbConnection grobal
var DbConnection *sql.DB

func GetDb() (*sql.DB, error) {
	DbConnection, err := sql.Open("mysql", "db_username:db_password@tcp(localhost:3306)/table_name")
	if err != nil {
		panic(err)
	} else {
		log.Println("-------------------->Connection DB OK")
	}

	return DbConnection, err
}