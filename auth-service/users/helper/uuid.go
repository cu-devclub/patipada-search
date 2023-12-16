package helper

import "github.com/google/uuid"

// GenerateUUID generates a UUID (Universally Unique Identifier).
//
// It takes no parameters.
// It returns a string representation of the generated UUID and an error, if any.
func GenerateUUID() (string, error) {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return newUUID.String(), nil
}
