package regex

import (
	"regexp"
)

/// regex email valid
func RegexEmail(nameregex string) bool {
	// regex for my endpoint
	var loginRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`) // Contains email valid
	if loginRegex.MatchString(nameregex) {
		return true
	} else {
		return false
	}
}
