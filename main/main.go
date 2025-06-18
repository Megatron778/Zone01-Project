package main

import (
	"cohort"
	"fmt"
	"strings"
)

func main() {

	text := "(cap) was the best of times, it was the worst of times (up) ,it was the age of wiSdOm (low)"
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

	for i := 0; i < len(slice); i++ {
		if i > 0 && slice[i] == "(up)" {
			slice[i-1] = strings.ToUpper(slice[i-1])
			slice[i] = ""
			slice = CleanSlice(slice)
		} else if i > 0 && slice[i] == "(low)" {
			slice[i-1] = strings.ToLower(slice[i-1])
			slice[i] = ""
			slice = CleanSlice(slice)
		} else if  i > 0 && slice[i] == "(cap)" {
			slice[i-1] = goreloaded.Capitalize(slice[i-1])
			slice[i] = ""
			slice = CleanSlice(slice)
    	}
   }

	return slice
}