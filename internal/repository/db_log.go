package repository

import (
	"database/sql"
	"log"

	"github.com/alejandro-rl/gamelogger-backend/internal/domain"
)

func CreateLog(db *sql.DB, log_var *domain.Log) error {
	query := `
	INSERT INTO log
	(replay,plat_id,game_id,user_id,status_id)
	VALUES (?,?,?,?,?)
	`

	//Insert user into table
	_, err := db.Exec(query, log_var.Replay, log_var.PlatID, log_var.GameID, log_var.UserID, log_var.StatusID)

	if err != nil {
		log.Print("Could not insert log into table")
		log.Print(err.Error())
		return err
	}

	return nil

}

func GetLogByID(db *sql.DB, log_id int) (*domain.Log, error) {
	query := `
	SELECT * FROM log where 
	log_id = ?
	`
	row := db.QueryRow(query, log_id)
	var log_var *domain.Log
	err := row.Scan(&log_var.ID, &log_var.Replay, &log_var.PlatID, &log_var.GameID, &log_var.UserID, &log_var.StatusID)

	if err != nil {
		log.Print("could not query log from table")
		log.Print(err.Error())
		return nil, err
	}

	return log_var, nil

}

func GetLogByUserGameID(db *sql.DB, user_id int, game_id int) (*domain.Log, error) {
	query := `
	SELECT * FROM log where 
	user_id = ? AND game_id = ?
	`
	row := db.QueryRow(query, user_id, game_id)
	var log_var *domain.Log
	err := row.Scan(&log_var.ID, &log_var.Replay, &log_var.PlatID, &log_var.GameID, &log_var.UserID, &log_var.StatusID)

	if err != nil {
		log.Print("could not query log from table")
		log.Print(err.Error())
		return nil, err
	}

	return log_var, nil

}

func UpdateLog(db *sql.DB, log_var *domain.Log) error {

	query := "UPDATE log SET replay = ?, plat_id = ?,game_id = ?,status_id = ? WHERE log_id = ?"
	_, err := db.Exec(query, log_var.Replay, log_var.PlatID, log_var.GameID, log_var.StatusID, log_var.ID)

	if err != nil {
		log.Print("Could not update log")
		log.Print(err.Error())
		return err
	}

	return nil

}

func DeleteLog(db *sql.DB, log_id int) error {

	query := "DELETE FROM log  WHERE log_id = ?"
	_, err := db.Exec(query, log_id)

	if err != nil {
		log.Print("Could not delete log")
		return err
	}

	return nil

}
