package models

import "time"

type CardStatus struct {
	CardID    string    `bson:"card_id" json:"card_id"`
	UserID    string    `bson:"user_id" json:"user_id"`
	Status    string    `bson:"status" json:"status"`
	Comment   string    `bson:"comment" json:"comment"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
}
