package goreloaded

func SolvePunc(st string) string {

	s := []rune(st)
	str := ""

	coutnb := 0
	count := 0

	for i := 0; i < len(st); i++ {
		if i != 0 && i != len(s)-1 && (s[i] == '\'' && ((IsPunctuation(s[i-1]) || IsPunctuation(s[i+1]) || s[i-1] == ' ' ||
			s[i+1] == ' ') || s[i+1] == '\'' || s[i-1] == '\'' || s[i-1] == '(' || s[i-1] == ')' )) {
			coutnb++
		} else if s[i] == '\'' && (i == 0 || i == len(s)-2) {
			coutnb++
		}
	}

	if coutnb%2 == 0 {
		coutnb = -1
	}

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
			s[i+1] == ' ') || s[i+1] == '\'' || s[i-1] == '\'' || s[i-1] == '(' || s[i-1] == ')' )) && count+1 != coutnb {
			str += " " + string(s[i]) + " "
			count++
		} else if s[i] == '\'' && (i == 0 || i == len(s)-2) && count+1 != coutnb {
			str += " " + string(s[i]) + " "
			count++
		} else {
			str += string(s[i])
		}
	}

	return str
}