package repository

import (
	"database/sql"

	"github.com/alejandro-rl/gamelogger-backend/internal/domain"
)

func CreateGenre(db *sql.DB, genre *domain.Genre) error {
	query := `
	INSERT INTO genre
	(igdb_id,genre)
	VALUES (?,?)
	`
	_, err := db.Exec(query, genre.IgdbID, genre.Name)

	if err != nil {
		return err
	}

	return nil

}
