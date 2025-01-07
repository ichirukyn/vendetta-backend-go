package entities

type Class struct {
	ID        string `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	DescShort string `json:"desc_short"`
	RaceID    string `json:"race_id"`
	Type      string `json:"type"`
	MainAttr  string `json:"main_attr"`
	Hidden    bool   `json:"hidden"`
	TagID     int    `json:"tag_id"`

	Race Race `json:"race,omitempty"`
	Tag  Tag  `json:"tag,omitempty"`
}

type ClassEffect struct {
	ID        string `json:"id" gorm:"primaryKey;autoIncrement"`
	ClassId   string `json:"class_id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Attribute string `json:"attribute"`
	Value     string `json:"value"`
}
