package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/alejandro-rl/gamelogger-backend/internal/domain"
	"github.com/alejandro-rl/gamelogger-backend/internal/repository"
	"github.com/go-chi/jwtauth/v5"
	"github.com/gorilla/mux"
)

func createUserHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var user_req *domain.UserRequest
		json.NewDecoder(r.Body).Decode(&user_req)

		err := repository.CreateUser(db, user_req)

		if err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "User created successfully")

	}
}

func getUserHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get the 'id' parameter from the URL
		vars := mux.Vars(r)
		id := vars["id"]
		user_id, err := strconv.Atoi(id)

		if err != nil {
			http.Error(w, "Please provide the correct input!", http.StatusBadRequest)
			return
		}

		//fetch user data
		user, err := repository.GetUserByID(db, user_id)

		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		//User object to JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)

	}
}

func updateUserHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get the 'id' parameter from the URL
		vars := mux.Vars(r)
		id := vars["id"]
		user_id, err := strconv.Atoi(id)

		if err != nil {
			http.Error(w, "Please provide the correct input!", http.StatusBadRequest)
			return
		}

		//First, verify if the user is authenticated
		if !VerifyAuth(r, user_id) {
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
			return
		}

		//If autheticated, user data can  be modified
		var user *domain.User
		json.NewDecoder(r.Body).Decode(&user)

		err = repository.UpdateUser(db, user)

		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		fmt.Fprintln(w, "User updated successfully")

	}
}

func deleteUserHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get the 'id' parameter from the URL
		vars := mux.Vars(r)
		id := vars["id"]
		user_id, err := strconv.Atoi(id)

		if err != nil {
			http.Error(w, "Please provide the correct input!", http.StatusBadRequest)
			return
		}

		//First, verify if the user is authenticated
		if !VerifyAuth(r, user_id) {
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
			return
		}

		//If autheticated, user data can  be deleted
		err = repository.DeleteUser(db, user_id)

		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		fmt.Fprintln(w, "User deleted successfully")

	}
}

func loginUserHandler(db *sql.DB, AuthToken *jwtauth.JWTAuth) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var user_req *domain.UserRequest
		json.NewDecoder(r.Body).Decode(&user_req)

		//See if the hashed password stored in the database is equal to the one provided
		tokenString, err := repository.LoginUser(db, AuthToken, user_req)

		if err != nil {
			http.Error(w, "Login failed - Please try again.", http.StatusBadRequest)
			return
		}

		//Send access token
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(tokenString))

	}
}

func VerifyAuth(r *http.Request, user_id int) bool {
	/* After the Verifier and Authenticator have successful validated this request
	* We destructure the claims from the request and get the userId from claims
	* We then check whether the userId from claims is same as the userId for which
	* the request has been hit (from url params), if not that means user is using
	* different JWT token and hence unauthorized.
	 */
	_, claims, _ := jwtauth.FromContext(r.Context())
	userIdFromClaims := int(claims["id"].(float64))

	if user_id != userIdFromClaims {
		return false
	} else {
		return true
	}
}
