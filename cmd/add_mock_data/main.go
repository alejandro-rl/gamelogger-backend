package main

import (
	"database/sql"
	"log"
	"os"

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

}
