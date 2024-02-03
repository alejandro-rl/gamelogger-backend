package repository

import (
	"database/sql"

	"github.com/alejandro-rl/gamelogger-backend/internal/domain"
)

func CreateGame(db *sql.DB, game *domain.Game) error {
	query := `
	INSERT INTO game
	(igdb_id,name,release_date,description)
	VALUES (?,?,?,?)
	`
	_, err := db.Exec(query, game.IgdbID, game.Name, game.ReleaseDate, game.Description)

	if err != nil {
		return err
	}

	return nil

}
