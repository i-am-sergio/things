package models

import (
	"fmt"
	"time"
)

type Notification struct {
	UserID    string    `json:"userId" bson:"userId"`
	Title     string    `json:"title" bson:"title"`
	Message   string    `json:"message" bson:"message"`
	Image     string    `json:"image" bson:"image"`
	IsRead    bool      `json:"isRead" bson:"isRead"`
	Type      string    `json:"type" bson:"type"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

func (n *Notification) TestFunction() {
	fmt.Println("This is a test function")
}
