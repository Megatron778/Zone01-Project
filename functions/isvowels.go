package goreloaded

import "strings"

func IsVowels(s string) bool {
	
	return strings.Contains("aeuiohAEUIHO", s)
}