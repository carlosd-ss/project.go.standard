// Go Api server
// @jeffotoni
// 2019-01-21

package util

import (
	"errors"
	"regexp"
	"strings"
)

// substr (text, 2, 30)
func Substr(value string, leni, lenf int) string {
	if len(value) < leni {
		return ""
	}

	lenx := len(value) // max of text
	leny := lenf       // amount of character

	if lenx < lenf { // amount of character
		leny = lenx
	}

	// return..
	return value[leni:leny]
}

// convert the first character of each word to upper case
func Ucwords(str string) string {

	// Function replacing words (assuming lower case input)
	replace := func(word string) string {
		return strings.Title(word)
	}

	r := regexp.MustCompile(`\w+`)
	str = r.ReplaceAllStringFunc(strings.ToLower(str), replace)

	return str
}

func Separate(str, separator string) (string, string, error) {
	sepIndex := strings.Index(str, separator)
	if sepIndex >= 0 {
		return str[0:(sepIndex)], str[sepIndex+len(separator):], nil
	} else {
		return "", "", errors.New("Separator now found!")
	}
}

func Partition(s string, sep string) (string, string, error) {
	parts := strings.SplitN(s, sep, 2)
	if len(parts) == 1 {
		return parts[0], "", nil
	}
	return parts[0], parts[1], nil
}
