package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	Username     string             `bson:"username"`
	Email        string             `bson:"email"`
	FirstName    string             `bson:"firstName"`
	LastName     string             `bson:"lastName"`
	Role         string             `bson:"role"`
	HashPassword string             `bson:"hashPassword"`
	CreatedAt    time.Time          `bson:"createdAt,omitempty"`
	UpdatedAt    time.Time          `bson:"updatedAt,omitempty"`
}

type NewUser struct {
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required, email"`
	FirstName string `json:"firstName" validate:"required"` 
	LastName  string `json:"lastName" validate:"required"`
	Role      string `json:"role" validate:"required"`
	Password  string `json:"password" validate:"required"`
}
