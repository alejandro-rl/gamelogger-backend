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
