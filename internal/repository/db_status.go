package repository

import (
	"database/sql"

	"github.com/alejandro-rl/gamelogger-backend/internal/domain"
)

func CreateStatus(db *sql.DB, status *domain.Status) error {
	query := `
	INSERT INTO status
	(status_id,status_name)
	VALUES (?,?)
	`
	_, err := db.Exec(query, status.ID, status.Name)

	if err != nil {
		return err
	}

	return nil

}

func GetStatusByID(db *sql.DB, status_id int) (*domain.Status, error) {
	query := `
	SELECT * FROM status WHERE status_id = ?
	`
	row := db.QueryRow(query, status_id)
	status := &domain.Status{}
	err := row.Scan(&status.ID, &status.Name)

	if err != nil {
		return nil, err
	}

	return status, nil

}

func GetStatusByName(db *sql.DB, status_name string) (*domain.Status, error) {
	query := `
	SELECT * FROM status WHERE status_name = ?
	`
	row := db.QueryRow(query, status_name)
	status := &domain.Status{}
	err := row.Scan(&status.ID, &status.Name)

	if err != nil {
		return nil, err
	}

	return status, nil

}
