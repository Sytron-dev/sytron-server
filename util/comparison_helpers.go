package util

import "github.com/google/uuid"

func IsEmptyUUID(u uuid.UUID) bool {
	return u == uuid.UUID{}
}
