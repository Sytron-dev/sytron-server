package models

import (
	"github.com/google/uuid"
)

type Company struct {
	SqlDocument `json:",inline" db:",inline"`

	// relations
	HQ uuid.UUID `json:"_hq" db:"_hq"`

	Name    string `json:"name"     db:"name"`
	Email   string `json:"email"    db:"email"`
	Phone   string `json:"phone"    db:"phone"`
	LogoURL string `json:"logo_url" db:"logo_url"`
}
