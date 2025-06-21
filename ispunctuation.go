package goreloaded

func IsPunctuation(r rune) bool {
	flag := false

	if r == '.' || r == ',' || r == '!' || r == ':' || r == ';' || r == '?' {
		flag = true
	}

	return flag
}
