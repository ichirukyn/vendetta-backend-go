package entities

import (
	"time"
)

type User struct {
	ID        string    `json:"id" gorm:"primaryKey;autoIncrement"`
	ChatID    string    `json:"chat_id"`
	Login     string    `json:"login"`
	Role      string    `json:"role"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserBan struct {
	ID     string `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID string `json:"user_id"`
	BanID  string `json:"ban_id"`
}
