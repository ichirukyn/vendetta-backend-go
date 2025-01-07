package dto

// CreateHeroDTO дто для создания класса
type CreateHeroDTO struct {
	UserID   string  `json:"user_id"`
	RaceID   string  `json:"race_id"`
	ClassID  string  `json:"class_id"`
	Name     string  `json:"name"`
	Rank     int     `json:"rank"`
	IsHidden bool    `json:"is_hidden"`
	Level    int     `json:"level"`
	Exp      float64 `json:"exp"`
}

// UpdateHeroDTO дто для обновления класса
type UpdateHeroDTO struct {
	Name     string  `json:"name"`
	Rank     int     `json:"rank"`
	IsHidden bool    `json:"is_hidden"`
	Level    int     `json:"level"`
	Exp      float64 `json:"exp"`
}

// FindHeroesDTO дто для поиска классов
type FindHeroesDTO struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Name   string `json:"name"`
}
