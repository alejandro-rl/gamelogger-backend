package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/alejandro-rl/gamelogger-backend/internal/domain"
	"github.com/alejandro-rl/gamelogger-backend/internal/repository"
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

	// Add game to database

	var game []domain.Game

	err = json.Unmarshal(igdb(), &game)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", game)

	err = repository.CreateGame(db, &game[0])

	if err != nil {
		log.Fatal(err)
	}

	/*
		//Create tables on the database
		path := filepath.Join("db/schema.sql")

		c, ioErr := os.ReadFile(path)
		if ioErr != nil {
			log.Fatal(err)
		}

		query := string(c)

		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	*/

}

func igdb() []byte {

	url := "https://api.igdb.com/v4/games"
	method := "POST"

	payload := strings.NewReader(`fields name,first_release_date,summary,genres,platforms; 
  	where id = 1942;
  	`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	req.Header.Add("Client-ID", os.Getenv("IGDBID"))
	req.Header.Add("Authorization", "Bearer "+os.Getenv("IGDBAUTH"))
	req.Header.Add("Content-Type", "text/plain")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	fmt.Println(string(body))
	return body
}
