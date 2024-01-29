package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {

	// Capture connection properties.
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "172.17.0.2",
		DBName:               "gamelogger",
		AllowNativePasswords: true,
		ParseTime:            true,
		MultiStatements:      true,
	}

	// Opens the connection

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	//Verifies the connnection

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(err)
	}

	//Create tables
	create_tables(db)

}

func create_tables(db *sql.DB) {

	//Create tables on the database

	path := filepath.Join("../../db/schema.sql")

	c, ioErr := os.ReadFile(path)
	if ioErr != nil {
		log.Fatal(ioErr)
	}

	query := string(c)

	_, err := db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}
}
