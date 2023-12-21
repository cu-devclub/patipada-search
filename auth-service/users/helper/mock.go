package helper

import "math/rand"

func GenerateRandomUsername() string {
	// Define a list of characters to choose from
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Generate a random username of length 10
	username := make([]byte, 10)
	for i := 0; i < 10; i++ {
		username[i] = chars[rand.Intn(len(chars))]
	}

	return string(username)
}

func GenerateRandomEmail() string {
	return GenerateRandomUsername() + "@example.com"
}
