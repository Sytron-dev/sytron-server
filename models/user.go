package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `json:"id"`
	First_name string    `json:"first_name"`
	Last_name  string    `json:"last_name"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserCredential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
