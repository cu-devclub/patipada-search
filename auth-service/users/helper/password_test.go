package helper

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestGenerateHashedSaltedPassword(t *testing.T) {
	rawPassword := "password123"

	hashedPassword, salt, err := GenerateHashedSaltedPassword(rawPassword)
	if err != nil {
		t.Errorf("Error generating hashed and salted password: %v", err)
	}

	// Verify that the hashed password is not empty
	if len(hashedPassword) == 0 {
		t.Errorf("Generated hashed password is empty")
	}

	// Verify that the salt is not empty
	if len(salt) == 0 {
		t.Errorf("Generated salt is empty")
	}

	// Verify that the hashed password is valid
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword+salt))
	if err != nil {
		t.Errorf("Generated hashed password is not valid: %v", err)
	}
}
