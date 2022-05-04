package validator

import (
	"regexp"
)

func init() {
	reAscii = regexp.MustCompile("^[a-zA-Z0-9_-]*$")
}

var reAscii *regexp.Regexp

// AsciiValidator returns true if input is [a-z0-9_-]
func AsciiValidator(ascii string) bool {
	return reAscii.MatchString(ascii)
}
