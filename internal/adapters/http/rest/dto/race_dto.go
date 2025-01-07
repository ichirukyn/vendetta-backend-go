package dto

// FindRacesDTO дто для поиска расы
type FindRacesDTO struct {
	ID     string `json:"id"`
	Hidden bool   `json:"hidden"`
}

// FindRacesEffectsDTO дто для поиска эффектов расы
type FindRacesEffectsDTO struct {
	RaceID string `json:"race_id"`
}
