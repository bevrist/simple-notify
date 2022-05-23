// validator implements data validation
package common

import (
	"regexp"
)

func init() {
	reAscii = regexp.MustCompile("^[a-zA-Z0-9_-]*$")
}

var reAscii *regexp.Regexp

// asciiValidator returns true if input is [a-z0-9_-]
func asciiValidator(ascii string) bool {
	return reAscii.MatchString(ascii)
}

func (m Message) Validate() error {
	return nil // TODO: implement this validator
}
