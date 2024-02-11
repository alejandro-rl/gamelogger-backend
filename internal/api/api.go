package api

import (
	"database/sql"
	"path/filepath"

	"github.com/gorilla/mux"
)

func Routes(db *sql.DB) *mux.Router {

	game_image_path := filepath.Join("../../db/images")

	//New router
	r := mux.NewRouter()

	//Game Routes
	r.HandleFunc("/game", createGameHandler(db, game_image_path)).Methods("POST")

	//Genre Routes
	r.HandleFunc("/genre", createGenreHandler(db)).Methods("POST")

	//Platform Routes
	r.HandleFunc("/platform", createPlatformHandler(db)).Methods("POST")

	return r

}
