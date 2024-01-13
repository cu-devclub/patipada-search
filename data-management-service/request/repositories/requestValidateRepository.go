package repositories

//TODO : Generate GoDoc
func (r *requestRepositories) ValidateRecordIndex(recordID string) (bool,error) {
	result,err := r.communicationClient.SearchRecord(recordID)
	if err != nil {
		return false,err
	}
	if result == false {
		return false,nil
	}
	
	return true,nil 
}


// ValidateUsername checks if the provided username is valid.
// It uses the communication client's VerifyUsername method to perform the check.
// If an error occurs during the operation, it will be returned along with a false boolean.
//
// Parameters:
//   username: The username to validate.
//
// Returns:
//   bool: A boolean indicating whether the username is valid. True if the username is valid, false otherwise.
//   error: An error that occurred during the operation, if any.
func (r *requestRepositories) ValidateUsername(username string) (bool, error) {
	result, err := r.communicationClient.VerifyUsername(username)
	if err != nil {
		return false, err
	}
	if result == false {
		return false, nil
	}

	return true, nil
}