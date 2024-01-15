package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `bson:"username,omitempty" binding:"required"`
	Email     string             `bson:"email,omitempty" binding:"required"`
	CreatedAt time.Time          `bson:"createdAt,omitempty" binding:"required"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty" binding:"required"`
}

type Admins struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" validate:"required"`
	Email     string             `json:"email,omitempty" validate:"required"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
}
