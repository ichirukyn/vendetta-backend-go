package entities

type Skill struct {
	ID        string  `json:"id" gorm:"primaryKey;autoIncrement"`
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

//type SkillEffect struct {
//	ID        string `json:"id" gorm:"primaryKey;autoIncrement"`
//	SkillId   string `json:"skill_id"`
//	Name      string `json:"name"`
//	Type      string `json:"type"`
//	Attribute string `json:"attribute"`
//	Value     string `json:"value"`
//}
