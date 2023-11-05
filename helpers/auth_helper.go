package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword is used to encrypt the password before it is stored in the DB
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// returns true if password is correct
func VerifyPassword(hashedPassword, password string) (isValid bool) {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
