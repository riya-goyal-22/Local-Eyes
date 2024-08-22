package utils

func ValidateUsername(username string) bool {
	return len(username) > 0
}

func ValidatePassword(password string) bool {
	return len(password) > 6
}

func ValidateUserType(userType string) bool {
	return userType == "admin" || userType == "newbie" || userType == "resident"
}

func GenerateID() string {
	// Simulate ID generation
	return "ID_" + randomString(10)
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[i%len(charset)]
	}
	return string(b)
}
