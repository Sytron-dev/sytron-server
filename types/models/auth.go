package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthCredential struct {
	SqlDocument `db:"sql_document"`

	Value        string             `json:"credential_value,omitempty" db:"credential_value"`
	Type         string             `json:"credential_type,omitempty"  db:"credential_type"` // email, phone, sso, username
	Password     string             `json:"password,omitempty"         db:"password"`
	BearerToken  string             `json:"bearer_token,omitempty"`
	RefreshToken string             `json:"refresh_token,omitempty"`
	LastLogin    primitive.DateTime `json:"last_login,omitempty"       db:"last_login"`
}

type AuthBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (doc *AuthCredential) UpdateLastLogin() {
	doc.LastLogin = primitive.NewDateTimeFromTime(time.Now())
}
