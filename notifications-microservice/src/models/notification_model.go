package models

import (
	"fmt"
	"time"
)

type NotificationModel struct {
	Id        string    `json:"_id" bson:"_id"`
	UserID    string    `json:"userId" bson:"userId"`
	Title     string    `json:"title" bson:"title"`
	Message   string    `json:"message" bson:"message"`
	Image     string    `json:"image" bson:"image"`
	IsRead    bool      `json:"isRead" bson:"isRead"`
	Type      string    `json:"type" bson:"type"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

func (n *NotificationModel) TestFunction() {
	fmt.Println("This is a test function")
}
