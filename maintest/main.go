package main

import (
	"fmt"
	"strconv"
	"strings"

	goreloaded "cohort"
)

func main() {
	text1 := "..(cap,5)!:;   .h!:;"

	text1 = strings.ToLower(text1)

	fmt.Println(text1)

	text2 := goreloaded.CleanStr(SolvePunc(text1))
	text3 := CleanPunc(goreloaded.StringToSlice(text2))
	fmt.Println(text3)
	text4 := GoReloaded(text3)
	fmt.Println(text4)
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

func CleanPunc(slice []string) []string {
	for i := 0; i < len(slice); i++ {
		if i != 0 && goreloaded.IsPunctuation(rune(slice[i][0])) && slice[i-1][len(slice[i-1])-1] != ')' {
			slice[i-1] += slice[i]
			slice[i] = ""
			slice = CleanSlice(slice)
			i--
		}
	}
	return slice
}

func SolvePunc(st string) string {
	s := []rune(st)
	str := ""

	for i := 0; i < len(s); i++ {
		if goreloaded.IsPunctuation(s[i]) {
			if i != 0 && i != len(s)-1 && !goreloaded.IsPunctuation(s[i-1]) && !goreloaded.IsPunctuation(s[i+1]) {
				str += " " + string(s[i]) + " "
			} else if i != 0 && !goreloaded.IsPunctuation(s[i-1]) {
				str += " " + string(s[i])
			} else if i != len(s)-1 && !goreloaded.IsPunctuation(s[i+1]) {
				str += string(s[i]) + " "
			} else {
				str += string(s[i])
			}
		} else {
			str += string(s[i])
		}
	}
	return str
}

func GoReloaded(slice []string) []string {
	number := 0

	for i := 0; i < len(slice); i++ {

		if slice[0] == "(up)" || slice[0] == "(low)" || slice[0] == "(cap)" || slice[0] == "(hex)" || slice[0] == "(bin)" {
			slice[0] = ""
			slice = CleanSlice(slice)
		} else if slice[0] == "(up," || slice[0] == "(low," || slice[0] == "(cap," {
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
		} else if i != 0 && slice[i] == "(cap)" {
			slice[i-1] = goreloaded.Capitalize(slice[i-1])
			slice[i] = ""
			slice = CleanSlice(slice)
			i--
		} else if i != 0 && slice[i] == "(hex)" {
			num, err := strconv.ParseInt(slice[i-1], 16, 64)

			if err != nil {
				fmt.Printf("Error converting '%s': %v\n", slice[i-1], err)
			} else {
				slice[i-1] = strconv.Itoa(int(num))
				slice[i] = ""
				slice = CleanSlice(slice)
				i--
			}
		} else if i != 0 && slice[i] == "(bin)" {
			num, err := strconv.ParseInt(slice[i-1], 2, 64)

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
			for j := 1; j <= number; j++ {
				if i-j >= 0 {
					slice[i-j] = strings.ToUpper(slice[i-j])
				}
			}
			slice[i] = ""
			slice[i+1] = ""
			slice = CleanSlice(slice)
			i--
		} else if i != 0 && i < len(slice)-1 && slice[i] == "(low," {
			number, _ = strconv.Atoi(slice[i+1][:len(slice[i+1])-1])
			for j := 1; j <= number; j++ {
				if i-j >= 0 {
					slice[i-j] = strings.ToLower(slice[i-j])
				}
			}
			slice[i] = ""
			slice[i+1] = ""
			slice = CleanSlice(slice)
			i--
		} else if i != 0 && i < len(slice)-1 && slice[i] == "(cap," {
			number, _ = strconv.Atoi(slice[i+1][:len(slice[i+1])-1])
			for j := 1; j <= number; j++ {
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
}
