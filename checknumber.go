package goreloaded

func CheckNumber(nb string) bool {
	flag := true

	star := 0

	if nb[0] == '-' || nb[0] == '+' {
		star = 1
	}

	for i := star ; i < len(nb); i++ {
		if nb[i] < '0' || nb[i] > '9' {
			flag = false
		}
	}
	return flag
}
