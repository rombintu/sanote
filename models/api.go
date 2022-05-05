package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type REST struct {
	Message string `json:"message"`
	Error   int    `json:"error"`
}

type Model struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	DeletedAt time.Time          `bson:"deleted_at" json:"deleted_at"`
}
