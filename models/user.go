package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	ID        primitive.ObjectID `bson:"_id"`
	UserName  string             `json:"username"`
	FirstName string             `json:"firstName"`
	LastName  string             `json:"lastName"`
	Password  string             `json:"password"`
	Status    string             `json:"status"`
	Email     string             `json:"email"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at"`
}
