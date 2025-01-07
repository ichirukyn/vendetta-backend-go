package entities

type Tag struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Desc         string `json:"desc"`
	Priority     int    `json:"priority"`
	Strength     int    `json:"strength"`
	Health       int    `json:"health"`
	Speed        int    `json:"speed"`
	Dexterity    int    `json:"dexterity"`
	Accuracy     int    `json:"accuracy"`
	Intelligence int    `json:"intelligence"`
	Submission   int    `json:"submission"`
	Soul         int    `json:"soul"`
}
