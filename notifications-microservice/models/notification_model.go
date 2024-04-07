package models

import "time"

type Notification struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	Message   string    `json:"message"`
	IsRead    bool      `json:"isRead"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
