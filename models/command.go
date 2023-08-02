package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommandModel struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `json:"name"`
	Detail    string             `json:"detail"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at"`
}
