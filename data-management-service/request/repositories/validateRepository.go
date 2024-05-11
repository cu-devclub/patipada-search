package repositories

// ValidateRecordIndex checks if a record with the given recordID exists.
// It communicates with the external service using the communicationClient to search for the record.
// If an error occurs during the operation, the function returns false and the error.
// If the record does not exist, the function returns false and nil.
// If the record exists, the function returns true and nil.
//
// Parameters:
//
//	recordID: The ID of the record to validate.
//
// Returns:
//
//	bool: A boolean indicating whether the record exists.
//	error: An error that occurred during the operation, if any.
func (r *repositoryImpl) ValidateRecordIndex(recordID string) (bool, error) {
	if recordID == "" {
		return false, nil
	}
	result, err := r.communicationClient.SearchRecord(recordID)
	if err != nil {
		return false, err
	}
	if !result {
		return false, nil
	}

	return true, nil
}

// ValidateUsername checks if the provided username is valid.
// It uses the communication client's VerifyUsername method to perform the check.
// If an error occurs during the operation, it will be returned along with a false boolean.
//
// Parameters:
//
//	username: The username to validate.
//
// Returns:
//
//	bool: A boolean indicating whether the username is valid. True if the username is valid, false otherwise.
//	error: An error that occurred during the operation, if any.
func (r *repositoryImpl) ValidateUsername(username string) (bool, error) {
	result, err := r.communicationClient.VerifyUsername(username)
	if err != nil {
		return false, err
	}
	if !result {
		return false, nil
	}

	return true, nil
}
