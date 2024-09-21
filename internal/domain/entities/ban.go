package entities

type Ban struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Reason    string `json:"reason"`
	CreatedAt string `json:"created_at"`
	Until     string `json:"until"`
}
