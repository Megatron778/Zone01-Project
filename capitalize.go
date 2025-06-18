package goreloaded

import "strings"

func Capitalize(s string) string {
	slice := []rune(s)
	str := ""
	
	for i := 0; i < len(slice); i++ {
		if i == 0 {
			str += string(strings.ToUpper(string(slice[i])))
		} else {
			str += string(strings.ToLower(string(slice[i])))
		}
	}
	return str
}