package model

// Category struct
type Category struct {
	ID        int64  `json:"id"`
	Topic     string `json:"topic"`
	Level     int64  `json:"level"`
	LevelID   int64  `json:"level_id"`
	Status    string `json:"status"`
	Position  int64  `json:"position"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
