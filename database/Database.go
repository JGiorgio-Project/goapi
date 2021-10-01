package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//DbConnection global
var DbConnection *sql.DB

func GetDb() (*sql.DB, error) {
	//Connection properties.
	const user = "***"
	const password = "***"
	const dbName = "***"

	//Make DNS
	dns := user + ":" + password + "@tcp(localhost:3306)/" + dbName

	//Get a database handle.
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
