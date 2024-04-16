package helper

import (
	"encoding/hex"
	"testing"
	"time"
)

func TestGenerateResetToken(t *testing.T) {
	token := GenerateResetToken()

	// Verify that the generated token is not empty
	if len(token) == 0 {
		t.Errorf("Generated reset token is empty")
	}

	// Verify that the generated token is a valid hexadecimal string
	_, err := hex.DecodeString(token)
	if err != nil {
		t.Errorf("Generated reset token is not a valid hexadecimal string")
	}
}

func TestGenerateResetTokenExpiration(t *testing.T) {
	expiration := GenerateResetTokenExpiration()

	// Verify that the generated expiration time is in the future
	if expiration.Before(time.Now()) {
		t.Errorf("Generated reset token expiration time is not in the future")
	}
}
