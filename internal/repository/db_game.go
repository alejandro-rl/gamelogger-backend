package repository

import (
	"database/sql"
	"log"

	"github.com/alejandro-rl/gamelogger-backend/internal/domain"
)

func CreateGame(db *sql.DB, game *domain.Game) error {

	// Create game without genres and platforms since these are stored in other tables
	query := `
	INSERT INTO game
	(igdb_id,name,release_date,description,url_name)
	VALUES (?,?,?,?,?)
	`
	result, err := db.Exec(query, game.IgdbID, game.Name, game.ReleaseDate, game.Description, game.URL)

	if err != nil {
		log.Print("Could not execute query to create game")
		return err
	}

	// Use the ID of the added game to associate the game_genres and game_platforms tables
	added_game_id, err := result.LastInsertId()

	if err != nil {
		log.Print("Could not get the id of last inserted game")
		return err
	}

	game.ID = int(added_game_id)

	//Associate Game and Genres

	err = GameGenres(db, game)

	if err != nil {
		return err
	}

	return nil

}

func GetGameByID(db *sql.DB, id int) (*domain.Game, error) {
	query := `
	SELECT * FROM game WHERE game_id = ?
	`
	row := db.QueryRow(query, id)
	game := &domain.Game{}
	err := row.Scan(&game.ID, &game.IgdbID, &game.Name, &game.ReleaseDate, &game.Description, &game.URL, &game.AverageRating)

	if err != nil {
		return nil, err
	}

	return game, nil

}

func GameGenres(db *sql.DB, game *domain.Game) error {
	query := `
	INSERT INTO game_genre
	(game_id,genre_id)
	VALUES (?,?)
	`

	// Iterate over genre list of a game
	for i := 0; i < len(game.Genres); i++ {
		genre, err := GetGenreByIgdbID(db, game.Genres[i])

		if err != nil {
			return err
		}

		//Insert game_id and genre_id into table

		_, err = db.Exec(query, game.ID, genre.ID)

		if err != nil {
			log.Print("Could not Insert values into game_genre table")
			return err
		}
	}

	return nil

}
