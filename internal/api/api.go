package api

import (
	"database/sql"
	"path/filepath"

	"github.com/go-chi/jwtauth/v5"
	"github.com/gorilla/mux"
)

func Routes(db *sql.DB, AuthToken *jwtauth.JWTAuth) *mux.Router {

	game_image_path := filepath.Join("../../db/game_images")

	//New router
	r := mux.NewRouter()

	//Game Routes
	r.HandleFunc("/game", createGameHandler(db, game_image_path)).Methods("POST")
	r.HandleFunc("/game/{url_name}", getGameHandler(db)).Methods("GET")
	r.HandleFunc("/game_images/{id}", getGameImageHandler(db)).Methods("GET")

	//Genre Routes
	r.HandleFunc("/genre", createGenreHandler(db)).Methods("POST")

	//Platform Routes
	r.HandleFunc("/platform", createPlatformHandler(db)).Methods("POST")

	//User Routes
	r.HandleFunc("/register", createUserHandler(db)).Methods("POST")
	r.HandleFunc("/login", loginUserHandler(db, AuthToken)).Methods("POST")
	r.HandleFunc("/user/{id}", getUserHandler(db)).Methods("GET")
	r.HandleFunc("/user/{id}", updateUserHandler(db)).Methods("PUT")
	r.HandleFunc("/user/{id}", deleteUserHandler(db)).Methods("DELETE")

	return r

}
