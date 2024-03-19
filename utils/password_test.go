package utils

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	password := RandomString(6)
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		t.Errorf("Hashed password does not match original password")
	}

	anatherPassword := RandomString(6)
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(anatherPassword))
	if err == nil {
		t.Errorf("Hashed password should not match another password")
	}
}