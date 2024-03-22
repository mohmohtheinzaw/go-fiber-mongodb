package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" `
	Username  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Admins struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" `
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
