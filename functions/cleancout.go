package goreloaded


func CleanCout(slice []string, flag bool) []string {
	count := 0

	coutnb := 0

	for i := 0; i < len(slice); i++ {
		if slice[i] == "'" {
			coutnb++
		} 
	}

	if coutnb%2 == 0 {
		coutnb = -1
	}
	for i := 0; i < len(slice); i++ {

		if len(slice[i]) != 0 && slice[i] != "'" && flag {

			if slice[i][0] == '\'' && len(slice[i]) != 1 {
				count++
				
			}
			if slice[i][len(slice[i])-1] == '\'' {
				count++
				
			}
		}


	    if i != len(slice)-1 && slice[i] == "'" && count%2 == 0 && !ValidFlag2(slice[i+1]) && count+1 != coutnb {
			if slice[i+1] == "'" {
				count++
			}
			slice[i+1] = slice[i] + slice[i+1]
			slice[i] = ""
			slice = CleanSlice(slice)
			i--
			count++
		} else if i != 0 && slice[i] == "'" && count%2 != 0 && !ValidFlag2(slice[i-1]) && count+1 != coutnb {
			if slice[i-1] == "'" {
				count++
			}
			slice[i-1] += slice[i]
			slice[i] = ""
			slice = CleanSlice(slice)
			i--
			count++
			
		 } else if i < len(slice)-2 && slice[i] == "'" && count%2 == 0 && !ValidFlag2(slice[i+1]) && !ValidFlag3(slice[i+1], slice[i+2]) && count+1 != coutnb {
			if slice[i+1] == "'" {
		 		count++
		 	}
		 	slice[i+1] = slice[i] + slice[i+1]
			slice[i] = ""
		 	slice = CleanSlice(slice)
		 	i--
		 	count++
		 }else if i > 1 && i < len(slice)-1 && slice[i] == "'" && count%2 == 0 && !ValidFlag2(slice[i-1]) && !ValidFlag3(slice[i-2], slice[i-1]) && count+1 != coutnb {
			if slice[i+1] == "'" {
		 		count++
		 	}
		 	slice[i+1] = slice[i] + slice[i+1]
		 	slice[i] = ""
		 	slice = CleanSlice(slice)
		 	i--
			count++
		} else if i != 0 && IsPunctuation(rune(slice[i][0])) && slice[i-1][len(slice[i-1])-1] != ')' {
			slice[i-1] += slice[i]
			slice[i] = ""
			slice = CleanSlice(slice)
			i--
		} 
	}
	return slice
}