package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthCredential struct {
	CollectionDocument `bson:",inline"`

	Value        string             `bson:"value"         json:"value"`
	Type         string             `bson:"type"          json:"type"` // email, phone, sso, username
	Password     string             `bson:"password"      json:"password"`
	BearerToken  string             `bson:"bearer_token"  json:"bearer_token"`
	RefreshToken string             `bson:"refresh_token" json:"refresh_token"`
	LastLogin    primitive.DateTime `bson:"last_login"    json:"last_login"`
}

func (doc *AuthCredential) UpdateLastLogin() {
	doc.LastLogin = primitive.NewDateTimeFromTime(time.Now())
}
