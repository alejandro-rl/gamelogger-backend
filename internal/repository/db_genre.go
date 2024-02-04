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

func GetGenreByIgdbID(db *sql.DB, igdb_id int) (*domain.Genre, error) {
	query := `
	SELECT * FROM genre WHERE igdb_id = ?
	`
	row := db.QueryRow(query, igdb_id)
	genre := &domain.Genre{}
	err := row.Scan(&genre.ID, &genre.IgdbID, &genre.Name)

	if err != nil {
		return nil, err
	}

	return genre, nil

}

func GetGenreByID(db *sql.DB, genre_id int) (*domain.Genre, error) {
	query := `
	SELECT * FROM genre WHERE genre_id = ?
	`
	row := db.QueryRow(query, genre_id)
	genre := &domain.Genre{}
	err := row.Scan(&genre.ID, &genre.IgdbID, &genre.Name)

	if err != nil {
		return nil, err
	}

	return genre, nil

}
