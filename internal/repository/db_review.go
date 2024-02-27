package repository

import (
	"database/sql"
	"log"

	"github.com/alejandro-rl/gamelogger-backend/internal/domain"
)

func CreateReview(db *sql.DB, review *domain.Review) error {
	query := `
	INSERT INTO review
	(score,favorite,time_played_total,review_text,log_id)
	VALUES (?,?,?,?,?)
	`

	//Insert user into table
	_, err := db.Exec(query, review.Score, review.Favorite, review.TimePlayedTotal, review.ReviewText, review.LogID)

	if err != nil {
		log.Print("Could not insert review into table")
		log.Print(err.Error())
		return err
	}

	return nil

}

func GetReviewByID(db *sql.DB, review_id int) (*domain.Review, error) {
	query := `
	SELECT * FROM review where 
	review_id = ?
	`
	row := db.QueryRow(query, review_id)
	var review *domain.Review
	err := row.Scan(&review.ID, &review.Score, &review.Favorite, &review.TimePlayedTotal,
		&review.TotalLikes, &review.TotalComments, &review.ReviewText, &review.LogID)

	if err != nil {
		log.Print("could not query review from table")
		log.Print(err.Error())
		return nil, err
	}

	return review, nil

}

func GetReviewByLogID(db *sql.DB, log_id int) (*domain.Review, error) {
	query := `
	SELECT * FROM review where 
	log_id = ?
	`
	row := db.QueryRow(query, log_id)
	var review *domain.Review
	err := row.Scan(&review.ID, &review.Score, &review.Favorite, &review.TimePlayedTotal,
		&review.TotalLikes, &review.TotalComments, &review.ReviewText, &review.LogID)

	if err != nil {
		log.Print("could not query review from table")
		log.Print(err.Error())
		return nil, err
	}

	return review, nil
}

func UpdateReview(db *sql.DB, review *domain.Review) error {

	query := "UPDATE review SET score = ?, favorite = ?,game_id = ?,time_played_total = ? WHERE log_id = ?"
	_, err := db.Exec(query, review.Score, review.Favorite, review.TimePlayedTotal, review.ReviewText, review.ID)

	if err != nil {
		log.Print("Could not update review")
		log.Print(err.Error())
		return err
	}

	return nil

}

func DeleteReview(db *sql.DB, review_id int) error {

	query := "DELETE FROM review  WHERE review_id = ?"
	_, err := db.Exec(query, review_id)

	if err != nil {
		log.Print("Could not delete review")
		return err
	}

	return nil

}
