package main

import (
	goreloaded "cohort"
	"fmt"
	"strings"
)

func main() {
	text := "..(cap,5)!:;   .h!:;"

	text = strings.ToLower(text)

	fmt.Println(text)
	
	text2 := goreloaded.CleanStr(Punctuatuion(text))
	text3 := goreloaded.StringToSlice(text2)
	fmt.Println(text3)
	fmt.Println(goreloaded.CleanPunc(text3))

	
}

func Punctuatuion(s string) string {

	str := ""
	count := 0

	for i := 0; i < len(s); i++ {
		if s[i] == ',' || s[i] == '.' || s[i] == ';' || s[i] == ':' || s[i] == '!' || s[i] == '?' {
			if i != 0 && (s[i-1] != ',' && s[i-1] != '.' && s[i-1] != ';' && s[i-1] != ':' && s[i-1] != '!' && s[i-1] != '?')  {
			str += " " + string(s[i])
			count = 1
			if i != len(s)-1 && (s[i+1] != ',' && s[i+1] != '.' && s[i+1] != ';' &&	s[i+1] != ':' && s[i+1] != '!' && s[i+1] != '?') {
				if count == 1 {
					str += " "
				} else {
					str +=  string(s[i]) + " "
				}
				count = 0
		    }
		} else if i != len(s)-1 && (s[i+1] != ',' && s[i+1] != '.' && s[i+1] != ';' &&	s[i+1] != ':' && s[i+1] != '!' && s[i+1] != '?') {
				
					str +=  string(s[i]) + " "
				
		    }else {
			str += string(s[i]) + " "
		}
		} else {
			str += string(s[i])
		}
	}

	return str
}