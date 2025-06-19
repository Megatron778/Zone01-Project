package main

import (
	"cohort"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	text := ",the best of times (cap) , it was  the wor,st  of  éimes (cap, 2)  it was tHe AGe of wiSdOm (cap, 2) ,"
	str := "" 

	fmt.Println(text)

	strclean := goreloaded.CleanStr(text)

	

	slice := (goreloaded.StringToSlice(strclean))

	fmt.Println(slice)

	finalslice := GoReloaded(slice)

	fmt.Println(finalslice)

	for i := 0; i < len(finalslice); i++ {
		if i != len(finalslice) {
			str += string(finalslice[i]) + " "
		} else {
			str += string(finalslice[i])
		}
	}
	fmt.Println(str)

 
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

func GoReloaded(slice []string) []string {

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
		} else if i != 0 && slice[i] == "(low)" {
			slice[i-1] = strings.ToLower(slice[i-1])
			slice[i] = ""
			slice = CleanSlice(slice)
		} else if  i != 0 && slice[i] == "(cap)" {
			slice[i-1] = goreloaded.Capitalize(slice[i-1])
			slice[i] = ""
			slice = CleanSlice(slice)
    	} else if i != 0 && slice[i] == "(hex)" || slice[i] == "(bin)" {
			if slice[i] == "(hex)" {
				num , err := strconv.ParseInt(slice[i-1], 16 , 64)

				if err != nil {
					fmt.Printf("Error converting '%s': %v\n", slice[i-1], err)
				} else {
					slice[i-1] = strconv.Itoa(int(num))
					slice[i] = ""
			        slice = CleanSlice(slice)
				}
			} else if slice[i] == "(bin)" {
				num2 , err2 := strconv.ParseInt(slice[i-1], 2 , 64)

				if err2 != nil {
					fmt.Printf("Error converting '%s': %v\n", slice[i-1], err2)
				} else {
					slice[i-1] = strconv.Itoa(int(num2))
					slice[i] = ""
			        slice = CleanSlice(slice)
				}
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
		}
   }

        if slice[0][0] == ',' || slice[0][0] == '.' || slice[0][0] == ':' || slice[0][0] == ';' || slice[0][0] == '?' || slice[0][0] == '!' {
			slice[0] = string(slice[0][0]) + " " + slice[0][1:]
		} 
    for i := 0; i < len(slice); i++ {
		
		for j := 0; j < len(slice[i]); j++ {
		    if j != 0 && j != len(slice[i])-1 && slice[i][j] == ',' || slice[i][j] == '.' || slice[i][j] == ';' || slice[i][j] == ':' || slice[i][j] == '?' || slice[i][j] == '!' {
				slice[i] = slice[i][:j+1] + " " + slice[i][j+1:]
			}
 	    }
		if i != 0 && slice[i] == "," || slice[i] == "." || slice[i] == ";" || slice[i] == ":" || slice[i] == "!" || slice[i] == "?" {
			slice[i-1] += slice[i]
			slice[i] = ""
			slice = CleanSlice(slice)
		}else if i != 0 && slice[i][0] == ',' || slice[i][0] == '.' || slice[i][0] == ';' || slice[i][0] == ':' || slice[i][0] == '!' || slice[i][0] == '?' {
			slice[i-1] += string(slice[i][0])
			slice[i] = slice[i][1:]
		}
    }

	return slice
}