package dto

// CreateSkillDTO дто для создания
type CreateSkillDTO struct {
	Name      string  `json:"name"`
	Desc      string  `json:"desc"`
	DescShort string  `json:"desc_short"`
	Rank      int     `json:"rank"`
	Damage    float64 `json:"damage"`
	Element   string  `json:"element"`
	Distance  string  `json:"distance"`
	Type      string  `json:"type"`
	MainStat  string  `json:"main_stat"`
	Cooldown  int     `json:"cooldown"`
	UserID    string  `json:"user_id"`
	ClassID   string  `json:"class_id"`
	RaceID    string  `json:"race_id"`
	IsStack   bool    `json:"is_stack"`
	Hidden    bool    `json:"hidden"`
}

// UpdateSkillDTO дто для обновления
type UpdateSkillDTO struct {
	Name      string  `json:"name"`
	Desc      string  `json:"desc"`
	DescShort string  `json:"desc_short"`
	Rank      int     `json:"rank"`
	Damage    float64 `json:"damage"`
	Element   string  `json:"element"`
	Distance  string  `json:"distance"`
	Type      string  `json:"type"`
	MainStat  string  `json:"main_stat"`
	Cooldown  int     `json:"cooldown"`
	ClassID   string  `json:"class_id"`
	RaceID    string  `json:"race_id"`
	IsStack   bool    `json:"is_stack"`
	Hidden    bool    `json:"hidden"`
}

// FindSkillsDTO дто для поиска
type FindSkillsDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Rank     int    `json:"rank"`
	Element  string `json:"element"`
	Distance string `json:"distance"`
	Type     string `json:"type"`
	MainStat string `json:"main_stat"`
	UserID   string `json:"user_id"`
	ClassID  string `json:"class_id"`
	RaceID   string `json:"race_id"`
	IsStack  bool   `json:"is_stack"`
	Hidden   bool   `json:"hidden"`
}
