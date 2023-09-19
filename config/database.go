package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var (
	db  *sql.DB
	err error
)

func handleDatabaseConnection() {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "root"
	dbname := "hacktiv8_assignment2"
	connectionString := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"

	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Panic("Error connecting to database", err)
	}

	err = db.Ping()
	if err != nil {
		log.Panic("Error connecting to database", err)
	}
}

func GetDatabaseConnection() *sql.DB {
	if db == nil {
		handleDatabaseConnection()
	}

	return db
}
