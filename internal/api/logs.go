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

func createLogHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var log_var *domain.Log
		json.NewDecoder(r.Body).Decode(&log_var)

		// Get the 'id' parameter from the URL
		vars := mux.Vars(r)
		u_id := vars["user_id"]

		// Convert 'id' to an integer
		user_id, err := strconv.Atoi(u_id)

		if err != nil {
			http.Error(w, "Please provide the correct input!", http.StatusBadRequest)
			return
		}

		log_var.UserID = user_id

		err = repository.CreateLog(db, log_var)

		if err != nil {
			http.Error(w, "Failed to create log", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "log created successfully")

	}
}

func getLogHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get the 'id' parameter from the URL
		vars := mux.Vars(r)
		l_id := vars["log_id"]

		// Convert 'id' to an integer
		log_id, err := strconv.Atoi(l_id)

		if err != nil {
			http.Error(w, "Please provide the correct input!", http.StatusBadRequest)
			return
		}

		//fetch log data
		log_var, err := repository.GetLogByID(db, log_id)

		if err != nil {
			http.Error(w, "Log not found", http.StatusNotFound)
			return
		}

		//User object to JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(log_var)

	}
}

func updateLogHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get the 'id' parameter from the URL
		vars := mux.Vars(r)
		u_id := vars["user_id"]

		// Convert 'id' to an integer
		user_id, err := strconv.Atoi(u_id)

		if err != nil {
			http.Error(w, "Please provide the correct input!", http.StatusBadRequest)
			return
		}

		//First, verify if the user is authenticated
		if !VerifyAuth(r, user_id) {
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
			return
		}

		var log_var *domain.Log
		json.NewDecoder(r.Body).Decode(&log_var)

		err = repository.UpdateLog(db, log_var)

		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		fmt.Fprintln(w, "User updated successfully")

	}
}

func deleteLogHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get the 'id' parameter from the URL
		vars := mux.Vars(r)
		u_id := vars["user_id"]
		l_id := vars["log_id"]

		// Convert 'id' to an integer
		user_id, err := strconv.Atoi(u_id)

		if err != nil {
			http.Error(w, "Please provide the correct input!", http.StatusBadRequest)
			return
		}

		//First, verify if the user is authenticated
		if !VerifyAuth(r, user_id) {
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
			return
		}

		// Convert 'id' to an integer
		log_id, err := strconv.Atoi(l_id)

		if err != nil {
			http.Error(w, "Please provide the correct input!", http.StatusBadRequest)
			return
		}

		//If autheticated, user data can  be deleted
		err = repository.DeleteLog(db, log_id)

		if err != nil {
			http.Error(w, "Log not found", http.StatusNotFound)
			return
		}

		fmt.Fprintln(w, "Log deleted successfully")

	}
}
