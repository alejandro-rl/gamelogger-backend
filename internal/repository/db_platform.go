package repository

import (
	"database/sql"

	"github.com/alejandro-rl/gamelogger-backend/internal/domain"
)

func CreatePlatform(db *sql.DB, platform *domain.Platform) error {
	query := `
	INSERT INTO platform
	(igdb_id,platform)
	VALUES (?,?)
	`
	_, err := db.Exec(query, platform.IgdbID, platform.Name)

	if err != nil {
		return err
	}

	return nil

}

func GetPlatformByIgdbID(db *sql.DB, igdb_id int) (*domain.Platform, error) {
	query := `
	SELECT * FROM platform WHERE igdb_id = ?
	`
	row := db.QueryRow(query, igdb_id)
	platform := &domain.Platform{}
	err := row.Scan(&platform.ID, &platform.IgdbID, &platform.Name)

	if err != nil {
		return nil, err
	}

	return platform, nil

}
