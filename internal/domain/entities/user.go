package entities

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	ChatID    string    `json:"chat_id"`
	Login     string    `json:"login"`
	Role      string    `json:"role"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) GenerateID() {
	u.ID = uuid.NewV4().String()
}

type UserBan struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	BanID  string `json:"ban_id"`
}
