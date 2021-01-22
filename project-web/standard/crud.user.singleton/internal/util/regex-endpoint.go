// Go Api server
// @jeffotoni
// 2019-05-14

package util

import (
	"net/http"
	"strings"
)

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
