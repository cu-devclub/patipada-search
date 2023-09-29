package util

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	// Generate a new UUID
	id := uuid.New()

	// Convert the UUID to a string
	idString := id.String()

	return idString
}

func EscapeText(input string) string {
    //Replace new line charcter with white space
    s := strings.ReplaceAll(input, "\n", " ")

    //Replace tab character with white space
    s = strings.ReplaceAll(s,"\t"," ")

	// Replace double quotes (") with a special character (e.g., '@@') for escaping
    s = strings.ReplaceAll(s, `"`, `@@`)
	return s
}

func UnescapeDoubleQuotes(input string) string {
	// Replace the special character (e.g., '@@') with double quotes (") for unescaping
	return strings.ReplaceAll(input, `@@`, `"`)
}
