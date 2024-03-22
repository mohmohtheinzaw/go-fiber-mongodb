package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	Id        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Username  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Admins struct {
	Id        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
