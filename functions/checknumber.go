package goreloaded

func CheckNumber(nb string) bool {
	if len(nb) < 1 {
		return false
	}
	star := 0
	if (nb[0] == '-' || nb[0] == '+' ) && len(nb) > 1 {
		star = 1
	}
	for i := star; i < len(nb); i++ {
		if nb[i] < '0' || nb[i] > '9' {
			return false
		}
	}
	return true
}
