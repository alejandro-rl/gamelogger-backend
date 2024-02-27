package domain

import "time"

type Review struct {
	ID              int       `json:"id"`
	Score           int       `json:"score"`
	Favorite        bool      `json:"favorite"`
	TimePlayedTotal time.Time `json:"time_played_total"`
	ReviewText      string    `json:"review_text"`
	TotalLikes      int       `json:"total_likes"`
	TotalComments   int       `json:"total_comments"`
	LogID           int       `json:"log_id"`
}
