package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DbConnection global
var DbConnection *sql.DB

func GetDb() (*sql.DB, error) {
	DbConnection, err := sql.Open("mysql", "db_username:db_password@tcp(localhost:3306)/table_name")
	if err != nil {
		panic(err)
	}

	return DbConnection, err
}
