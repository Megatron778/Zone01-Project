package goreloaded

import "unicode"

func IsWord(str string) bool {

	for _, r := range str {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}