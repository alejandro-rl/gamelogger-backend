package domain

type Log struct {
	ID       int `json:"id"`
	Replay   int `json:"replay"`
	PlatID   int `json:"plat_id"`
	GameID   int `json:"game_id"`
	UserID   int `json:"user_id"`
	StatusID int `json:"status_id"`
}
