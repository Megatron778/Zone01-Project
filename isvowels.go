package goreloaded

func IsVowels(s string) bool {
	flag := false
			if (s[0] == 'a' || s[0] == 'e' || s[0] == 'u' || s[0] == 'i' || s[0] == 'o' || s[0] == 'h' || s[0] == 'A' || s[0] == 'E' ||
				s[0] == 'U' || s[0] == 'I' || s[0] == 'O' || s[0] == 'H') {
				flag = true
	}

	return flag
}
