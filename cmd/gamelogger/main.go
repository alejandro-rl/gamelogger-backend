package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/alejandro-rl/gamelogger-backend/internal/api"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

func main() {
	//Start database connection
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

	//Access Token Generation
	AuthToken := GenerateAuthToken()

	//Routing
	r := api.Routes(db, AuthToken)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	log.Println("Server listening on :8090")
	log.Fatal(http.ListenAndServe(":8090", handler))

}

func GenerateAuthToken() *jwtauth.JWTAuth {
	//Access Token
	var JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")
	tokenAuth := jwtauth.New("HS256", []byte(JWT_SECRET_KEY), nil)
	return tokenAuth
}
