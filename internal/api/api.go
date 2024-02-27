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
	r.HandleFunc("/user/{user_id}", getUserHandler(db)).Methods("GET")

	//For the next routes, authentication is required
	auth_r := r.PathPrefix("/").Subrouter()
	auth_r.Use(jwtauth.Verifier(AuthToken))
	auth_r.Use(jwtauth.Authenticator(AuthToken))

	auth_r.HandleFunc("/user/{user_id}", updateUserHandler(db)).Methods("PUT")
	auth_r.HandleFunc("/user/{user_id}", deleteUserHandler(db)).Methods("DELETE")

	//Log routes

	auth_r.HandleFunc("/user/{user_id}/log/", createLogHandler(db)).Methods("POST")
	auth_r.HandleFunc("/user/{user_id}/log/{log_id}", getLogHandler(db)).Methods("GET")
	auth_r.HandleFunc("/user/{user_id}/log/{log_id}", updateLogHandler(db)).Methods("PUT")
	auth_r.HandleFunc("/user/{user_id}/log/{log_id}", deleteLogHandler(db)).Methods("DELETE")

	return r

}
