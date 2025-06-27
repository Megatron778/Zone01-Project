package goreloaded

func CleanSlice(slice []string) []string {
	sliceclean := []string{}
	for i := 0; i < len(slice); i++ {
		if slice[i] != "" {
			sliceclean = append(sliceclean, slice[i])
		}
	}
	return sliceclean
}