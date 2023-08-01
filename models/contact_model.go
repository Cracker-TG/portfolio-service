package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ContactModel struct {
	ID        primitive.ObjectID `bson:"_id"`
	Subject   string             `json:"subject"`
	Email     string             `json:"email"`
	Message   string             `json:"message"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at"`
}
