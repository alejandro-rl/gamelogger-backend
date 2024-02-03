package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alejandro-rl/gamelogger-backend/internal/domain"
	"github.com/alejandro-rl/gamelogger-backend/internal/repository"
)

func createGenreHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var genre []domain.Genre
		json.NewDecoder(r.Body).Decode(&genre)

		err := repository.CreateGenre(db, &genre[0])

		if err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "User created successfully")

	}
}
