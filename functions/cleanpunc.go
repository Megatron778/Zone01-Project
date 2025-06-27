package goreloaded


func CleanPunc(slice []string) []string {
	for i := 0; i < len(slice); i++ {

	if i != 0 && IsPunctuation(rune(slice[i][0])) && slice[i-1][len(slice[i-1])-1] != ')' && slice[i-1] != "'" {
			slice[i-1] += slice[i]
			slice[i] = ""
			slice = CleanSlice(slice)
			i--
		} 
	}
	return slice
}