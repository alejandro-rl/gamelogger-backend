package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/alejandro-rl/gamelogger-backend/internal/api"
	"github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

func main() {
	var db *sql.DB

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

	//Routing
	r := api.Routes(db)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	log.Println("Server listening on :8090")
	log.Fatal(http.ListenAndServe(":8090", handler))

}
