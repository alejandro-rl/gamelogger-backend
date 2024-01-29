package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alejandro-rl/gamelogger-backend/internal/domain"
	"github.com/alejandro-rl/gamelogger-backend/internal/repository"
)

func createGameHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var game []domain.Game
		json.NewDecoder(r.Body).Decode(&game)

		err := repository.CreateGame(db, &game[0])

		if err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "User created successfully")

	}
}
