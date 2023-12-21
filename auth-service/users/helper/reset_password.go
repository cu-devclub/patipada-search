package helper

import (
    "crypto/rand"
    "encoding/hex"
    "time"
)

func GenerateResetToken() string {
    tokenBytes := make([]byte, 32)
    rand.Read(tokenBytes)
    return hex.EncodeToString(tokenBytes)
}

func GenerateResetTokenExpiration() time.Time {
    return time.Now().Add(15 * time.Minute)
}
