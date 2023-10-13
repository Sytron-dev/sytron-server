package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthCredential struct {
	CollectionDocument `bson:",inline"`

	Credential     string             `json:"credential"`
	CredentialType string             `json:"credential_type"` // email, phone, sso, username
	Password       string             `json:"password"`
	BearerToken    string             `json:"bearer_token"`
	RefreshToken   string             `json:"refresh_token"`
	LastLogin      primitive.DateTime `json:"last_login,omitempty"`
}

func (doc AuthCredential) UpdateLastLogin() {
	doc.LastLogin = primitive.NewDateTimeFromTime(time.Now())
}
