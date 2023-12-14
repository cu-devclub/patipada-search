package helper

// GetStringField returns the value of a string field from a document.
//
// Parameters:
// - doc: a map[string]interface{} representing the document.
// - fieldName: a string representing the name of the field.
//
// Return:
// - a string representing the value of the field. If the field is not found or is not a string, it returns an empty string.
func GetStringField(doc map[string]interface{}, fieldName string) string {
	if value, found := doc[fieldName].(string); found {
		return value
	}
	return ""
}