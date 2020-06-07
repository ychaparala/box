package helpers

import (
	"fmt"
	"regexp"
	"unicode/utf8"
)

// ValidateEmail returns boolean
func ValidateEmail(email string) bool {
	var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	// email addresses have a practical limit of 254 bytes
	if len(email) > 254 || !rxEmail.MatchString(email) {
		fmt.Println("error: " + email + " is not a valid email address")
		return false
	}
	return true
}

// ValidatePassword returns boolean
func ValidatePassword(password string) bool {
	l := utf8.RuneCountInString(password)
	if l < 8 || l > 50 {
		fmt.Println("error: password must be between 8 and 50 characters long")
		return false
	}
	return true
}
