package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

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

		//Getting the id parameter from the URL
		vars := mux.Vars(r)
		url_name := vars["url_name"]

		//fetch ugame data
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
