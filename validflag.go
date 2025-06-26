package goreloaded

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