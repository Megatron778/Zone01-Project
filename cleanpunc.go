package goreloaded

func CleanPunc(slice []string) []string {
	sliceclean := []string{}
	for i := 0; i < len(slice); i++ {
		if i != 0 && len(slice[i]) > 1 && (slice[i][0] == '.' || slice[i][0] == ',' || slice[i][0] == ';' || slice[i][0] == ':' || slice[i][0] == '!' || slice[i][0] == '?') && (slice[i-1][len(slice[i-1])-1] != ')') {
			slice[i-1] += slice[i]
			slice[i] = ""
			for i := 0; i < len(slice); i++ {
				if slice[i] != "" {
					sliceclean = append(sliceclean, slice[i])
				}
			}
		}
	}

	return sliceclean
}
