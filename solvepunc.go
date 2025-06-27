package goreloaded

func SolvePunc(st string) string {

	s := []rune(st)
	str := ""

	for i := 0; i < len(s); i++ {
		if IsPunctuation(s[i]) {
			if i != 0 && i != len(s)-1 && !IsPunctuation(s[i-1]) && !IsPunctuation(s[i+1]) {
				str += " " + string(s[i]) + " "
			} else if i != 0 && !IsPunctuation(s[i-1]) {
				str += " " + string(s[i])
			} else if i != len(s)-1 && !IsPunctuation(s[i+1]) {
				str += string(s[i]) + " "
			} else {
				str += string(s[i])
			}
		} else if i != 0 && s[i] == '(' && IsPunctuation(s[i-1]) {
			str += " " + string(s[i])
		} else if i != len(s)-1 && s[i] == ')' && IsPunctuation(s[i+1]) {
			str += string(s[i]) + " "
		} else if i != 0 && i != len(s)-1 && (s[i] == '\'' && ((IsPunctuation(s[i-1]) || IsPunctuation(s[i+1]) || s[i-1] == ' ' ||
			s[i+1] == ' ') || s[i+1] == '\'' || s[i-1] == '\'' || s[i-1] == '(' || s[i-1] == ')' )) {
			str += " " + string(s[i]) + " "
		} else if s[i] == '\'' && (i == 0 || i == len(s)-1) {
			str += " " + string(s[i]) + " "
		} else {
			str += string(s[i])
		}
	}

	return str
}