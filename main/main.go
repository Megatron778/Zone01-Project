package main

import (
	"cohort"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	text := "(cap, 5) best of times 'élHo (cap) , it was .:! the wo..rst  hj4 (hex) (bin)  of  éimes hello world' (cap, 3) it was ,tHe AGe of wiSdOm (cap, 2) .."

	fmt.Println(text)
	text = goreloaded.Punctuatuion(text)
	strclean := goreloaded.CleanStr(text)
    slice := (goreloaded.StringToSlice(strclean))

	fmt.Println(slice)

	finalslice := GoReloaded(slice)

	fmt.Println(finalslice)
 
}

func CleanSlice(slice []string) []string {

	sliceclean := []string{}
	for i := 0; i < len(slice); i++ {
		if slice[i] != "" {
			sliceclean = append(sliceclean, slice[i])
		}
	}
	return sliceclean
}

func GoReloaded(slice []string) string {

	number := 0

	for i := 0; i < len(slice); i++ {

		if slice[0] == "(up)" || slice[0] == "(low)" || slice[0] == "(cap)" || slice[0] == "(hex)" || slice[0] == "(bin)" {
			slice[0] = ""
			slice = CleanSlice(slice)
		} else if (slice[0] == "(up," || slice[0] == "(low," || slice[0] == "(cap,") {
			slice[0] = ""
			slice[1] = ""
			slice = CleanSlice(slice)
		}


		if i != 0 && slice[i] == "(up)" {
			slice[i-1] = strings.ToUpper(slice[i-1])
			slice[i] = ""
			slice = CleanSlice(slice)
			i--
		} else if i != 0 && slice[i] == "(low)" {
			slice[i-1] = strings.ToLower(slice[i-1])
			slice[i] = ""
			slice = CleanSlice(slice)
			i--
		} else if  i != 0 && slice[i] == "(cap)" {
			slice[i-1] = goreloaded.Capitalize(slice[i-1])
			slice[i] = ""
			slice = CleanSlice(slice)
			i--
    	} else if i != 0 && slice[i] == "(hex)" {
				num , err := strconv.ParseInt(slice[i-1], 16 , 64)

				if err != nil {
					fmt.Printf("Error converting '%s': %v\n", slice[i-1], err)
				} else {
					slice[i-1] = strconv.Itoa(int(num))
					slice[i] = ""
			        slice = CleanSlice(slice)
					i--
				}
		} else if i != 0 && slice[i] == "(bin)" {
				num , err := strconv.ParseInt(slice[i-1], 2, 64)

				if err != nil {
					fmt.Printf("Error converting '%s': %v\n", slice[i-1], err)
				} else {
					slice[i-1] = strconv.Itoa(int(num))
					slice[i] = ""
			        slice = CleanSlice(slice)
					i--
				}
			
		} else if i != 0 && i < len(slice)-1 && slice[i] == "(up," {
			number, _ = strconv.Atoi(slice[i+1][:len(slice[i+1])-1])
			for j := 1; j <= number ; j++ {
				if i-j >= 0 {
					slice[i-j] = strings.ToUpper(slice[i-j])
				}
			}	
			        slice[i] = ""
					slice[i+1] = ""
			        slice = CleanSlice(slice)
					i--
		}else if i != 0 && i < len(slice)-1 && slice[i] == "(low," {
			number, _ = strconv.Atoi(slice[i+1][:len(slice[i+1])-1])
			for j := 1; j <= number ; j++ {
				if i-j >= 0 {
					slice[i-j] = strings.ToLower(slice[i-j])
				}
			}	
			        slice[i] = ""
					slice[i+1] = ""
			        slice = CleanSlice(slice)
					i--
		}else if i != 0 && i < len(slice)-1 && slice[i] == "(cap," {
			number, _ = strconv.Atoi(slice[i+1][:len(slice[i+1])-1])
			for j := 1; j <= number ; j++ {
				if i-j >= 0 {
					slice[i-j] = goreloaded.Capitalize(slice[i-j])
				}
			}	
			        slice[i] = ""
					slice[i+1] = ""
			        slice = CleanSlice(slice)
					i--
		}
   }

        if slice[0][0] == ',' || slice[0][0] == '.' || slice[0][0] == ':' || slice[0][0] == ';' || slice[0][0] == '?' || slice[0][0] == '!' {
			slice[0] = string(slice[0][0]) + " " + slice[0][1:]
		} 

		
    for i := 0; i < len(slice); i++ {


		for j := 0; j < len(slice[i]); j++ {
			
		    if j != 0 && j != len(slice[i])-1 && (slice[i][j] == ',' || slice[i][j] == '.' || slice[i][j] == ';' || slice[i][j] == ':' || slice[i][j] == '?' || slice[i][j] == '!') &&
			(slice[i][j+1] != '.' && slice[i][j+1] != ';' &&  slice[i][j+1] != ':' && slice[i][j+1] != '!' && slice[i][j+1] != ',' && slice[i][j+1] != '?') { 
				slice[i] = slice[i][:j+1] + " " + slice[i][j+1:]
			} 
 	    }
		if i != 0 && slice[i] == "," || slice[i] == "." || slice[i] == ";" || slice[i] == ":" || slice[i] == "!" || slice[i] == "?" {
			slice[i-1] += slice[i]
			slice[i] = ""
			slice = CleanSlice(slice)
			i--
		}
		
		

    }

	fmt.Println(slice[0])
	
	
	

	return goreloaded.SliceToString(slice)
}