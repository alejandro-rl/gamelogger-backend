package domain

type Platform struct {
	ID int

	IgdbID int    `json:"id"`
	Name   string `json:"name"`
}
