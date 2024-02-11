package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"

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

	//Create image folder if it doesn't exist yet
	game_img_path := filepath.Join("../../db/game_images")
	err = os.MkdirAll(game_img_path, os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

	//Add Genres
	AddGenres(db, "../../db/genres.json")

	//Add Platforms
	AddPlatforms(db, "../../db/platforms.json")

	//Add Games
	AddGames(db, "../../db/mock_games.json", game_img_path)

}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func OpenJSON(path string) *os.File {
	jsonFile, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Println(err)
	}

	return jsonFile

}

func AddGenres(db *sql.DB, path string) {

	var genres []domain.Genre

	jsonFile := OpenJSON(path)

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		log.Println(err)
	}

	json.Unmarshal(byteValue, &genres)

	for i := 0; i < len(genres); i++ {

		err = repository.CreateGenre(db, &genres[i])
		if err != nil {
			log.Println(err)
		}
	}

	jsonFile.Close()
}

func AddPlatforms(db *sql.DB, path string) {

	var platforms []domain.Platform

	jsonFile := OpenJSON(path)

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		log.Println(err)
	}

	json.Unmarshal(byteValue, &platforms)

	for i := 0; i < len(platforms); i++ {

		err = repository.CreatePlatform(db, &platforms[i])
		if err != nil {
			log.Println(err)
		}
	}

	jsonFile.Close()
}

func AddGames(db *sql.DB, path string, game_img_path string) {

	var games []domain.GameSet

	jsonFile := OpenJSON(path)

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		log.Println(err)
	}

	json.Unmarshal(byteValue, &games)

	for i := 0; i < len(games); i++ {

		err = repository.CreateGame(db, &games[i], game_img_path)
		if err != nil {
			log.Println(err)
		}
	}

	jsonFile.Close()
}
