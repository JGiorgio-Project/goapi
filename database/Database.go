package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//DbConnection global
var DbConnection *sql.DB

func GetDb() (*sql.DB, error) {
	// Capture connection properties.
	dns := "db_user:db_password@tcp(localhost:3306)/db_name"

	// Get a database handle.
	var err error
	DbConnection, err = sql.Open("mysql", dns)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := DbConnection.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	return DbConnection, err
}
