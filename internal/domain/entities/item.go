package entities

type Item struct {
	ID          string `json:"id"`
	ClassID     string `json:"class_id"`
	ClassType   string `json:"class_type"`
	RaceID      string `json:"race_id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Value       string `json:"value"`
	Description string `json:"description"`
	Modifier    string `json:"modifier"`
}
