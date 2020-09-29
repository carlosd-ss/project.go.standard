// Go Api server
// @jeffotoni
// 2019-05-14

package util

import (
	"regexp"
)

// regex email valid
func RegexPhone(nameregex string) bool {
	// regex for my endpoint
	var phone = regexp.MustCompile(`^(?:(?:\+|00)?(55)\s?)?(?:\(?([1-9][0-9])\)?\s?)?(?:((?:9\d|[2-9])\d{3})\-?(\d{4}))$`)
	if phone.MatchString(nameregex) {
		return true
	} else {
		return false
	}
}
