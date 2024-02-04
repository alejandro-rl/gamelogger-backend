package repository

import (
	"database/sql"
	"log"

	"github.com/alejandro-rl/gamelogger-backend/internal/domain"
)

func CreateGame(db *sql.DB, game *domain.GameSet) error {

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
	game_id := int(added_game_id)

	if err != nil {
		log.Print("Could not get the id of last inserted game")
		return err
	}

	//Associate Game and Genres in game_genre table

	err = SetGameGenres(db, game_id, game.Genres)

	if err != nil {
		return err
	}

	//Associate Game and Platforms in game_platform table

	err = SetGamePlatforms(db, game_id, game.Platforms)

	if err != nil {
		return err
	}

	return nil

}

func GetGameByID(db *sql.DB, id int) (*domain.GameGet, error) {
	//Query game info
	query := `
	SELECT * FROM game WHERE game_id = ?
	`
	row := db.QueryRow(query, id)
	game := &domain.GameGet{}
	err := row.Scan(&game.ID, &game.IgdbID, &game.Name, &game.ReleaseDate, &game.Description, &game.URL, &game.AverageRating)

	if err != nil {
		return nil, err
	}

	//Fill game.Genres with genre names
	genres, err := GetGameGenres(db, game.ID)

	if err != nil {
		return nil, err
	}

	for i := 0; i < len(genres); i++ {
		game.Genres = append(game.Genres, genres[i].Name)
	}

	//Fill game.Platforms with platform names
	platforms, err := GetGamePlatforms(db, game.ID)

	if err != nil {
		return nil, err
	}

	for i := 0; i < len(platforms); i++ {
		game.Platforms = append(game.Platforms, platforms[i].Name)
	}

	return game, nil

}

func SetGameGenres(db *sql.DB, game_id int, genres []int) error {
	query := `
	INSERT INTO game_genre
	(game_id,genre_id)
	VALUES (?,?)
	`

	// Iterate over genre list of a game
	for i := 0; i < len(genres); i++ {
		genre, err := GetGenreByIgdbID(db, genres[i])

		if err != nil {
			return err
		}

		//Insert game_id and genre_id into table

		_, err = db.Exec(query, game_id, genre.ID)

		if err != nil {
			log.Print("Could not Insert values into game_genre table")
			return err
		}
	}

	return nil

}

func GetGameGenres(db *sql.DB, game_id int) ([]*domain.Genre, error) {

	//Query game_genre to get all the genre_ids of a game
	query := `
	SELECT genre_id FROM game_genre WHERE game_id = ?
	`
	var genre_ids []int

	rows, err := db.Query(query, game_id)

	if err != nil {
		log.Print("Could not query rows in game_genre")
		return nil, err
	}

	//Iterate over rows to get the ids
	var id int
	for rows.Next() {

		err = rows.Scan(&id)
		genre_ids = append(genre_ids, id)

		if err != nil {
			log.Print("Could not scan game_genre values")
			return nil, err
		}
	}

	//With the ids, query the genre table to get the genres

	var genre_list []*domain.Genre
	for i := 0; i < len(genre_ids); i++ {
		genre, err := GetGenreByID(db, genre_ids[i])

		if err != nil {
			return nil, err
		}

		genre_list = append(genre_list, genre)

	}

	return genre_list, nil
}

func SetGamePlatforms(db *sql.DB, game_id int, platforms []int) error {
	query := `
	INSERT INTO game_platform
	(game_id,plat_id)
	VALUES (?,?)
	`

	// Iterate over platform list of a game
	for i := 0; i < len(platforms); i++ {
		platform, err := GetPlatformByIgdbID(db, platforms[i])

		if err != nil {
			return err
		}

		//Insert game_id and platform_id into table

		_, err = db.Exec(query, game_id, platform.ID)

		if err != nil {
			log.Print("Could not Insert values into game_platform table")
			return err
		}
	}

	return nil

}

func GetGamePlatforms(db *sql.DB, game_id int) ([]*domain.Platform, error) {

	//Query game_platform to get all the plat_ids of a game
	query := `
	SELECT plat_id FROM game_platform WHERE game_id = ?
	`
	var plat_ids []int

	rows, err := db.Query(query, game_id)

	if err != nil {
		log.Print("Could not query rows in game_platform")
		return nil, err
	}

	//Iterate over rows to get the ids
	var id int
	for rows.Next() {

		err = rows.Scan(&id)
		plat_ids = append(plat_ids, id)

		if err != nil {
			log.Print("Could not scan game_platform values")
			return nil, err
		}
	}

	//With the ids, query the platform table to get the platforms

	var plat_list []*domain.Platform
	for i := 0; i < len(plat_ids); i++ {
		genre, err := GetPlatformByID(db, plat_ids[i])

		if err != nil {
			return nil, err
		}

		plat_list = append(plat_list, genre)

	}

	return plat_list, nil
}
