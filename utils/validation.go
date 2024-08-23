package utils

import "strings"

func ValidateUsername(username string) bool {
	return len(username) > 0
}

func ValidatePassword(password string) bool {
	if len(password) > 6 {
		if strings.Contains(password, "@") || strings.Contains(password, "#") || strings.Contains(password, "$") || strings.Contains(password, "%") || strings.Contains(password, "^") || strings.Contains(password, "*") {
			if strings.Contains(password, "1") || strings.Contains(password, "2") || strings.Contains(password, "3") || strings.Contains(password, "4") || strings.Contains(password, "5") || strings.Contains(password, "6") || strings.Contains(password, "7") || strings.Contains(password, "8") || strings.Contains(password, "9") || strings.Contains(password, "0") {
				return true
			}
		}
	}
	return false
}

func ValidateUserType(userType string) bool {
	return userType == "newbie" || userType == "resident"
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
