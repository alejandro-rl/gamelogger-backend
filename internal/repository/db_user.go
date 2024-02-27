package repository

import (
	"database/sql"
	"log"

	"github.com/alejandro-rl/gamelogger-backend/internal/domain"
	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(db *sql.DB, user_req *domain.UserRequest) error {
	query := `
	INSERT INTO user
	(email,username,hash)
	VALUES (?,?,?)
	`

	// hash password
	hash_pass, err := getHashPassword(user_req.Password)

	if err != nil {
		return err
	}

	//Insert user into table
	_, err = db.Exec(query, user_req.Email, user_req.Username, hash_pass)

	if err != nil {
		return err
	}

	return nil

}

func GetUserByID(db *sql.DB, user_id int) (*domain.User, error) {
	query := `
	SELECT * FROM user where 
	user_id = ?
	`
	row := db.QueryRow(query, user_id)
	var user *domain.User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Description, &user.Hash)

	if err != nil {
		log.Print("could not query user from table")
		return nil, err
	}

	return user, nil

}

func GetUserByUsername(db *sql.DB, username string) (*domain.User, error) {
	query := `
	SELECT * FROM user where 
	username = ?
	`
	row := db.QueryRow(query, username)
	var user *domain.User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Description, &user.Hash)

	if err != nil {
		log.Print("could not query user from table")
		return nil, err
	}

	return user, nil

}

func UpdateUser(db *sql.DB, user *domain.User) error {

	query := "UPDATE user SET email = ?, username = ?,description = ? WHERE user_id = ?"
	_, err := db.Exec(query, user.Email, user.Username, user.Description, user.ID)

	if err != nil {
		log.Print("Could not update user")
		log.Print(err.Error())
		return err
	}

	return nil

}

func DeleteUser(db *sql.DB, user_id int) error {

	query := "DELETE FROM user  WHERE user_id = ?"
	_, err := db.Exec(query, user_id)

	if err != nil {
		log.Print("Could not delete user")
		return err
	}

	return nil

}

func LoginUser(db *sql.DB, AuthToken *jwtauth.JWTAuth, user_req *domain.UserRequest) (string, error) {

	query := `
	SELECT * FROM user where email = ?
	`
	row := db.QueryRow(query, user_req.Email)
	var user domain.User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Description, &user.Hash)

	if err != nil {
		log.Print("Could not query user from table")
		log.Print(err.Error())
		return "", err
	}

	//Verify user

	if !checkPassword(user.Hash, user_req.Password) {
		log.Print("Incorrect password")
		return "", err
	}

	//After verifying user credentials, generate a access token
	claims := map[string]interface{}{"id": user.ID, "email": user.Email, "username": user.Username}
	_, tokenString, err := AuthToken.Encode(claims)

	if err != nil {
		log.Print("Could not generate user access token")
		return "", err
	}
	return tokenString, nil

}

func getHashPassword(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		log.Print("Could not hash the user password")
		return "", err
	}
	return string(hash), nil
}

func checkPassword(hash_pass, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash_pass), []byte(password))
	return err == nil
}
