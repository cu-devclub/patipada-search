package helper

import (
	"crypto/rand"
	"fmt"
)

// GenerateSalt generates a random salt of length 16 bytes.
//
// It returns the salt as a hexadecimal string and an error if there was any
// issue generating the salt.
func GenerateSalt() (string, error) {
	// Generate a random salt of length 16 bytes
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", salt), nil
}
