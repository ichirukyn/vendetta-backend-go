package dto

// CreateUserDTO дто для создания пользователя
type CreateUserDTO struct {
	ChatID string `json:"chat_id"`
	Login  string `json:"login"`
	Role   string `json:"role"`
}

// UpdateUserDTO дто для обновления пользователя
type UpdateUserDTO struct {
	ChatID string `json:"chat_id"`
	Login  string `json:"login"`
	Role   string `json:"role"`
}

// FindUsersDTO дто для поиска пользователей
type FindUsersDTO struct {
	ChatID string `json:"chat_id"`
	Login  string `json:"login"`
	Role   string `json:"role"`
}
