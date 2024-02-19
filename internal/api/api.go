package api

import (
	"database/sql"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

func Routes(db *sql.DB) *mux.Router {

	game_image_path := filepath.Join("../../db/game_images")

	//New router
	r := mux.NewRouter()

	//Game Routes
	r.HandleFunc("/game", createGameHandler(db, game_image_path)).Methods("POST")
	r.HandleFunc("/game/{url_name}", getGameHandler(db)).Methods("GET")
	//r.Handle("/game_images", http.StripPrefix("/game_images", http.FileServer(http.Dir(game_image_path))))
	r.HandleFunc("/game_images/{id}", getGameImageHandler(db)).Methods("GET")

	fileServer := http.FileServer(http.Dir(game_image_path))
	r.Handle("/game_images/", http.StripPrefix("/game_images", fileServer))

	//Genre Routes
	r.HandleFunc("/genre", createGenreHandler(db)).Methods("POST")

	//Platform Routes
	r.HandleFunc("/platform", createPlatformHandler(db)).Methods("POST")

	return r

}
