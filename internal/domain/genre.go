package domain

type Genre struct {
	ID int

	IgdbID int    `json:"id"`
	Name   string `json:"name"`
}
