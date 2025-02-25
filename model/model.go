package model

import (
	"time"
)

type User struct {
	UserID    string    `bson:"user_id" json:"user_id"` // Store UUID separately
	Name      string    `bson:"name" json:"name"`
	Number    string    `bson:"number" json:"number"`
	Email     string    `bson:"email" json:"email"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}
