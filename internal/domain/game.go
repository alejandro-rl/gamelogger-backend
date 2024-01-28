package domain

type Game struct {
	ID int

	IgdbID        int    `json:"id"`
	Name          string `json:"name"`
	ReleaseDate   int    `json:"first_release_date"`
	Description   string `json:"summary"`
	AverageRating float32
}
