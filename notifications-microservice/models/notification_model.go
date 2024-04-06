package models

type Notification struct {
	ID      string `json:"id"`
	UserID  string `json:"userId"`
	Message string `json:"message"`
	IsRead  bool   `json:"isRead"`
}
