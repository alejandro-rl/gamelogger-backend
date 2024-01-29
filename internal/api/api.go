package api

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func Routes(db *sql.DB) *mux.Router {

	//New router
	r := mux.NewRouter()

	//Game Routes
	r.HandleFunc("/game", createGameHandler(db)).Methods("POST")

	return r

}
