package util

import (
	"html"

	"github.com/google/uuid"
	"github.com/microcosm-cc/bluemonday"
)

func Contains(s string, arr []string) bool {
	for _, a := range arr {
		if a == s {
			return true
		}
	}
	return false
}

func DecodeHTMLText(encodedString string) string {
	return html.UnescapeString(encodedString)
}

func ExtractRawStringFromHTMLTags(s string) string {
	// decode again first
	s = DecodeHTMLText(s)

	// then sanitize
	p := bluemonday.StrictPolicy()
	return p.Sanitize(s)
}

func GenerateUUID() string {
	return uuid.New().String()
}
