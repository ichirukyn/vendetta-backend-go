package dto

// CreateClassDTO дто для создания класса
type CreateClassDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	DescShort string `json:"desc_short"`
	RaceID    string `json:"race_id"`
	MainAttr  string `json:"main_attr"`
	Type      string `json:"type"`
	Hidden    bool   `json:"hidden"`
	TagID     int    `json:"tag_id"`
}

// UpdateClassDTO дто для обновления класса
type UpdateClassDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	DescShort string `json:"desc_short"`
	RaceID    string `json:"race_id"`
	MainAttr  string `json:"main_attr"`
	Type      string `json:"type"`
	Hidden    bool   `json:"hidden"`
	TagID     int    `json:"tag_id"`
}

// FindClassesDTO дто для поиска классов
type FindClassesDTO struct {
	ID       string `json:"id"`
	RaceID   string `json:"race_id"`
	MainAttr string `json:"main_attr"`
	Type     string `json:"type"`
	Hidden   bool   `json:"hidden"`
	TagID    int    `json:"tag_id"`
}

// FindClassesEffectsDTO дто для поиска эффектов расы
type FindClassesEffectsDTO struct {
	ClassID string `json:"class_id"`
}
