package goreloaded

func SliceToString(s []string) string {

	str := ""

	for i := 0; i < len(s); i++ {
		if i != len(s) {
			str += string(s[i]) + " "
		} else {
			str += string(s[i])
		}
	}
	return str
}

