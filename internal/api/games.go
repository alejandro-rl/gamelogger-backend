package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/alejandro-rl/gamelogger-backend/internal/domain"
	"github.com/alejandro-rl/gamelogger-backend/internal/repository"
	"github.com/gorilla/mux"
)

func createGameHandler(db *sql.DB, game_img_path string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var game []domain.GameSet
		json.NewDecoder(r.Body).Decode(&game)

		err := repository.CreateGame(db, &game[0], game_img_path)

		if err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "User created successfully")

	}
}

func getGameHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get the 'url_name' parameter from the URL
		vars := mux.Vars(r)
		url_name := vars["url_name"]

		//fetch game data
		game, err := repository.GetGameByURLName(db, url_name)

		if err != nil {
			http.Error(w, "Game not found", http.StatusNotFound)
			return
		}

		//Game object to JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(game)

	}
}

func getGameImageHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get the 'id' parameter from the URL
		vars := mux.Vars(r)
		id := vars["id"]

		// Convert 'id' to an integer
		game_id, err := strconv.Atoi(id)

		if err != nil {
			http.Error(w, "Please provide the correct input!", http.StatusBadRequest)
			return
		}

		//fetch game image path
		path, err := repository.GetGameImages(db, game_id)

		if err != nil {
			http.Error(w, "Game Image not found", http.StatusNotFound)
			return
		}

		http.ServeFile(w, r, path)

	}
}
