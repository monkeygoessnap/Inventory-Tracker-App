package server

import (
	"regexp"
)

//checks whether input is allowed based on regexp expression
//sanitization for inputs up till emails, prevents js scripts
func inputCheck(args ...string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9-@./ ]*$")
	for _, v := range args {
		if !re.MatchString(v) {
			return false
		}
	}
	return true
}

// checks whether input is alphanum based on regexp expression
// sanitization for inputs
func isNum(args ...string) bool {
	re := regexp.MustCompile("^[0-9]*$")
	for _, v := range args {
		if !re.MatchString(v) {
			return false
		}
	}
	return true
}
