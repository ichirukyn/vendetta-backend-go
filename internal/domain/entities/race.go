package entities

type Race struct {
	ID        string `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	DescShort string `json:"desc_short"`
	Hidden    bool   `json:"hidden"`
	TagID     int    `json:"tag_id"`

	Tag        Tag        `json:"tag,omitempty"`
	RaceEffect RaceEffect `json:"effects,omitempty"`
}

type RaceEffect struct {
	ID        string `json:"id" gorm:"primaryKey;autoIncrement"`
	RaceId    string `json:"race_id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Attribute string `json:"attribute"`
	Value     string `json:"value"`
}
