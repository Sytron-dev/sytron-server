package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	First_name *string    `json:"first_name" validate:"required,min=2,max=100"`
	Last_name  *string    `json:"last_name" validate:"required,min=2,max=100"`
	Username   *string    `json:"username" validate:"required,min=2,max=100"`
	Email      *string    `json:"email" validate:"required,email"`
	Password   *string 	  `json:"Password" validate:"required,min=6""`

	Token         *string  `json:"token"`
    Refresh_token *string  `json:"refresh_token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User_id  string    `json:"user_id"`
}

// type UserCredential struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }
