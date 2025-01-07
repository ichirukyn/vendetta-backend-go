package dto

// CreateHeroSpecDTO дто для создания класса
type CreateHeroSpecDTO struct {
	HeroID         string  `json:"hero_id"`
	Accuracy       int     `json:"accuracy"`
	Strength       int     `json:"strength"`
	Health         int     `json:"health"`
	Speed          int     `json:"speed"`
	Dexterity      int     `json:"dexterity"`
	Soul           int     `json:"soul"`
	Intelligence   int     `json:"intelligence"`
	Submissions    int     `json:"submissions"`
	CriticalRate   float64 `json:"critical_rate"`
	CriticalDamage float64 `json:"critical_damage"`
	Resistance     float64 `json:"resistance"`
	TotalSpec      int     `json:"total_spec"`
	FreeSpec       int     `json:"free_spec"`
}

// UpdateHeroSpecDTO дто для обновления класса
type UpdateHeroSpecDTO struct {
	Accuracy       int     `json:"accuracy"`
	Strength       int     `json:"strength"`
	Health         int     `json:"health"`
	Speed          int     `json:"speed"`
	Dexterity      int     `json:"dexterity"`
	Soul           int     `json:"soul"`
	Intelligence   int     `json:"intelligence"`
	Submissions    int     `json:"submissions"`
	CriticalRate   float64 `json:"critical_rate"`
	CriticalDamage float64 `json:"critical_damage"`
	Resistance     float64 `json:"resistance"`
	TotalSpec      int     `json:"total_spec"`
	FreeSpec       int     `json:"free_spec"`
}

// FindHeroSpecDTO дто для поиска классов
type FindHeroSpecDTO struct {
	ID     string `json:"id"`
	HeroID string `json:"hero_id"`
}
