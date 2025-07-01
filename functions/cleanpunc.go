package goreloaded

func CleanPunc(slice []string) []string {

	for i := 0; i < len(slice); i++ {
		if i != 0 && IsPunctuation(rune(slice[i][0])) && slice[i-1][len(slice[i-1])-1] != ')' && slice[i-1] != "'" {
			slice[i-1] += slice[i]
			slice[i] = ""
			slice = CleanSlice(slice)
			i--	
		 } else if i > 1 && IsPunctuation(rune(slice[i][0])) && (((slice[i-2] != "(up," && slice[i-2] != "(low," && slice[i-2] != "(cap,") || (!CheckNumber(slice[i-1][:len(slice[i-1])-1]) && slice[i-1][len(slice[i-1])-1] == ')')) &&
		 slice[i-1] != "(up)" && slice[i-1] != "(cap)" && slice[i-1] != "(low)" && slice[i-1] != "(bin)" && (slice[i-2] != "(hex)") && slice[i-2] != "(up)" && slice[i-2] != "(cap)" && slice[i-2] != "(low)" && slice[i-2] != "(bin)" && slice[i-2] != "(hex)") {
			slice[i-1] += slice[i]
			slice[i] = ""
			slice = CleanSlice(slice)
			i--
		} else if i > 0 && IsPunctuation(rune(slice[i][0])) && (!CheckNumber(slice[i-1][:len(slice[i-1])-1]) && slice[i-1][len(slice[i-1])-1] == ')') &&
		 slice[i-1] != "(up)" && slice[i-1] != "(cap)" && slice[i-1] != "(low)" && slice[i-1] != "(bin)" && slice[i-1] != "(hex)" {
			slice[i-1] += slice[i]
			slice[i] = ""
			slice = CleanSlice(slice)
			i--
		} else if i == 1 && IsPunctuation(rune(slice[i][0])) && (CheckNumber(slice[i-1][:len(slice[i-1])-1]) && slice[i-1][len(slice[i-1])-1] == ')') {
			slice[i-1] += slice[i]
			slice[i] = ""
			slice = CleanSlice(slice)
			i--
		}
		
		
	}
	return slice
}
