package domain

type Game struct {
	ID int

	IgdbID      int    `json:"id"`
	Name        string `json:"name"`
	ReleaseDate int    `json:"first_release_date"`
	Description string `json:"summary"`
	Genres      []int  `json:"genres"`
	Platforms   []int  `json:"platforms"`
	URL         string `json:"slug"`

	AverageRating float32
}
