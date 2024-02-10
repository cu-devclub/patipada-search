package helper

import (
	"search-esdb-service/record/entities"
	"strings"
)

// EscapeText replaces new line and tab characters with white space in the input string.
//
// Parameters:
// - input: the string to be escaped.
//
// Returns:
// - The escaped string.
func EscapeText(input string) string {
	//Replace new line charcter with white space
	s := strings.ReplaceAll(input, "\n", " ")

	//Replace tab character with white space
	s = strings.ReplaceAll(s, "\t", " ")

	// Replace double quotes (") with a special character (e.g., '@@') for escaping
	s = strings.ReplaceAll(s, `"`, `@@`)
	return s
}

// UnescapeDoubleQuotes replaces the special character (e.g., '@@') with double quotes (") for unescaping.
//
// It takes a single parameter:
// - input: a string that contains the special character to be replaced.
//
// It returns a string that has the special character replaced with double quotes.
func UnescapeDoubleQuotes(input string) string {
	// Replace the special character (e.g., '@@') with double quotes (") for unescaping
	return strings.ReplaceAll(input, `@@`, `"`)
}

func UnescapeFieldsAndCreateRecord(doc interface{}, docID string) *entities.Record {
	unescapedDoc := make(map[string]interface{})
	for key, value := range doc.(map[string]interface{}) {
		if stringValue, isString := value.(string); isString {
			// Unescape the string value
			unescapedValue := UnescapeDoubleQuotes(stringValue)
			unescapedDoc[key] = unescapedValue
		} else {
			unescapedDoc[key] = value
		}
	}
	unescapedDoc["id"] = docID

	return &entities.Record{
		Index:      docID,
		YoutubeURL: unescapedDoc["youtubeURL"].(string),
		Question:   unescapedDoc["question"].(string),
		Answer:     unescapedDoc["answer"].(string),
		StartTime:  unescapedDoc["startTime"].(string),
		EndTime:    unescapedDoc["endTime"].(string),
	}
}
