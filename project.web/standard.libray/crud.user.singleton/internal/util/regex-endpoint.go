// Go Api server
// @jeffotoni
// 2019-05-14

package util

import (
	"net/http"
	"regexp"
	"strings"
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

// regex email valid
func RegexEmail(nameregex string) bool {
	// regex for my endpoint
	var loginRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`) // Contains email valid
	if loginRegex.MatchString(nameregex) {
		return true
	} else {
		return false
	}
}

// generating endpoint and endpoint nameregex
func EndpointRegex(r *http.Request) (string, string) {
	vetPath := strings.Split(r.URL.Path, "/")
	lastpos := len(vetPath) - 1
	nameregex := vetPath[lastpos]
	endpoint := ""
	for _, v := range vetPath {
		if v != nameregex {
			endpoint += "/" + v
		}
	}
	// endpoint
	endpoint = "/" + strings.Trim(endpoint, "/")
	return endpoint, nameregex
}

// generating endpoint and endpoint nameregex
func EndpointRegexTwo(r *http.Request) (string, string) {
	vetPath := strings.Split(r.URL.Path, "/")
	lastpos := len(vetPath) - 2
	nameregex := vetPath[lastpos]
	endpoint := ""
	for _, v := range vetPath {
		if v != nameregex {
			endpoint += "/" + v
		}
	}
	// endpoint
	endpoint = "/" + strings.Trim(endpoint, "/")
	return endpoint, nameregex
}
