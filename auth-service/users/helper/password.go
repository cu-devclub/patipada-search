package helper

import "golang.org/x/crypto/bcrypt"

// GenerateHashedSaltedPassword generates a hashed and salted password.
//
// It takes a rawPassword string as input and returns a string representing
// the hashed password, a string representing the salt used, and an error if
// any occurred during the process.
//
// return hashedPassword, salt, err
func GenerateHashedSaltedPassword(rawPassword string) (string,string, error) {
	salt, err := GenerateSalt()
	if err != nil {
		return "","" ,err
	}
	// Combine password and salt
	saltedPassword := rawPassword + salt

	// Hash the salted password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), 10)
	if err != nil {
		return "","", err
	}
	return string(hashedPassword),salt, nil
}

// VerifyPassword verifies if a given password matches the hashed password stored in the database.
//
// Parameters:
// - dbPassword: the hashed password stored in the database.
// - inputPasswordAndSalt: the user's input password concatenated with the salt.
//
// Returns:
// - error: an error if the passwords do not match.
func VerifyPassword(dbPassword string, inputPasswordAndSalt string) error {
	return bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(inputPasswordAndSalt))
}
