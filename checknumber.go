package goreloaded

func CheckNumber(nb string) bool {
	flag := true 
	for _, v := range nb {
		if v < '0' || '9' > v {
			flag = false
		}
	}
	return flag
}
