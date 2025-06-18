package goreloaded

func CleanStr(s string) string {
	s += " "

	slice := []rune(s)
	var str string = ""
	var count int

	for i := 0; i < len(slice); i++ {
		if i != 0 && slice[i] != ' ' && slice[i-1] == ' ' {
			count = i
		} else if i != 0 && slice[i] == ' ' && slice[i-1] != ' ' {
			str += string(slice[count:i]) + " "
		}
	}

	return str[:len(str)-1]
}