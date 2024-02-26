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
	ID            int      `json:"id"`
	IgdbID        int      `json:"igdb_id"`
	Name          string   `json:"name"`
	ReleaseDate   int      `json:"release_date"`
	Description   string   `json:"description"`
	Genres        []string `json:"genres"`
	Platforms     []string `json:"platforms"`
	URL           string   `json:"url"`
	AverageRating float32  `json:"average_rating"`
}
