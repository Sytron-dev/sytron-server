package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	First_name   string             `json:"first_name" validate:"required,min=2,max=100"`
	Last_name    string             `json:"last_name" validate:"required,min=2,max=100"`
	Password     string             `bson:"password" validate:"required"`
	Email        string             `json:"email" validate:"email,required"`
	Phone        string             `json:"phone" validate:"required"`
	Token        string             `json:"token"`
	RefreshToken string             `json:"refresh_token"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
}

// type UserCredential struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }
