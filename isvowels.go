package goreloaded

import (
	"unicode"
)

func IsVowels(s string) bool {
	flag := false
	count := 0
	for i := 0; i < len(s); i++ {
		if unicode.IsLetter(rune(s[i])) {
			count++

			if (s[i] == 'a' || s[i] == 'e' || s[i] == 'u' || s[i] == 'i' || s[i] == 'o' || s[i] == 'h' || s[i] == 'A' || s[i] == 'E' ||
				s[i] == 'U' || s[i] == 'I' || s[i] == 'O' || s[i] == 'H') && count == 1 {
				flag = true
			}
		}
	}

	return flag
}
