package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	ID        primitive.ObjectID `bson:"_id"`
	UserName  string             `json:"username"`
	Name      string             `json:"prefix_name"`
	Password  string             `json:"password"`
	Active    bool               `json:"active,omitempty"`
	Email     string             `json:"email"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at"`
}
