package record

func GenerateUniqueDocuments(duplicateDocument []map[string]interface{})  []map[string]interface{} {
	// Create a map to track unique documents
	uniqueDocuments := make(map[string]map[string]interface{})

	// Filter duplicates and store unique documents in the map
	for _, doc := range duplicateDocument {
		// Use a unique identifier as the key, e.g., the document's ID or another field
		identifier := doc["id"].(string)
		uniqueDocuments[identifier] = doc
	}

	// Convert the map back to a slice
	uniqueDocumentSlice := make([]map[string]interface{}, 0, len(uniqueDocuments))
	for _, doc := range uniqueDocuments {
		uniqueDocumentSlice = append(uniqueDocumentSlice, doc)
	}
	return uniqueDocumentSlice
}
