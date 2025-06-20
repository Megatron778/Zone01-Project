package goreloaded

func Punctuatuion(s1 string) string {

	s := []string{}

	for i := 0; i < len(s1); i++ {
	    s = append(s, string(rune(s1[i])))
	}
	

	for i := 0; i < len(s); i++ {
		if i != len(s)-1 && (s[i] == "," || s[i] == "." || s[i] == ";" || s[i] == ":" || s[i] == "!" || s[i] == "?") &&  (s[i+1] != "," && s[i+1] != "." && s[i+1] != ";" && s[i+1] != ":" &&
		 s[i+1] != "!" && s[i+1] != "?") {
			s[i] += " "
			i--
		}
		if i != 0 && (s[i] == "," || s[i] == "." || s[i] == ";" || s[i] == ":" || s[i] == "!" || s[i] == "?") && (s[i-1] != "," && s[i-1] != "." && s[i-1] != ";" &&
			s[i-1] != ":" && s[i-1] != "!" && s[i-1] != "?") {
			s[i] = " " + (s[i])
			i--
		}
	}

	str := ""

	for i := 0; i < len(s); i++ {
			str += string(s[i])
	}

	return str
}