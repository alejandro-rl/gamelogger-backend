package api

import (
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {

	//New router
	r := mux.NewRouter()

	//Game Routes
	r.HandleFunc("/game", createGameHandler).Methods("POST")

	return r

}
