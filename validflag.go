package goreloaded

import "fmt"

func ValidFlag(s []string) bool {
	flag := false
	for i := 0 ; i < len(s) ; i++ {

		if s[i] == "(up)" || s[i] == "(low)" || s[i] == "(cap)" || s[i] == "(hex)" || s[i] == "(bin)" {
			flag = true
		}else if len(s) > 1 && i < len(s)-1 && ( s[i] == "(up," || s[i] == "(low," || s[i] == "(cap,") && CheckNumber(s[i+1][:len(s[i+1])-1]) && s[i+1][len(s[i+1])-1] == ')' {
			flag = true
			
		}
	}
	return flag
}

func ValidFlag2(s string) bool {

	flag := false
	if s == "(up)" || s == "(low)" || s == "(cap)" || s == "(hex)" || s == "(bin)" {
			flag = true
		}
	fmt.Println(flag)
	return flag
}

func ValidFlag3(s, s2 string) bool {

	flag := false
	
		if ( s == "(up," || s == "(low," || s == "(cap,") && CheckNumber(s2[:len(s2)-1]) && s2[len(s2)-1] == ')' {
			flag = true
		}
		fmt.Println(flag, "hh")
	
	return flag
}