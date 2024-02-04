package domain

// This struct handles game info coming from the IGDB API
type GameSet struct {
	IgdbID      int    `json:"id"`
	Name        string `json:"name"`
	ReleaseDate int    `json:"first_release_date"`
	Description string `json:"summary"`
	Genres      []int  `json:"genres"`
	Platforms   []int  `json:"platforms"`
	URL         string `json:"slug"`
}

// This struct handles game info coming from the database
type GameGet struct {
	ID            int
	IgdbID        int
	Name          string
	ReleaseDate   int
	Description   string
	Genres        []string
	Platforms     []string
	URL           string
	AverageRating float32
}
