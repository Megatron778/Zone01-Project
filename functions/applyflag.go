package goreloaded

import (
	"fmt"
	"strconv"
	"strings"
)

func ApplyFlag(slice []string) []string {
	number := 0
	for i := 0; i < len(slice); i++ {
		if (slice[0] == "(up)" || slice[0] == "(low)" || slice[0] == "(cap)" || slice[0] == "(hex)" || slice[0] == "(bin)") {
			slice[0] = ""
			slice = CleanSlice(slice)
		} else if len(slice) > 1 && (slice[0] == "(up," || slice[0] == "(low," || slice[0] == "(cap,") && CheckNumber(slice[1][:len(slice[1])-1]) && slice[1][len(slice[1])-1] == ')' {
			slice[0] = ""
			slice[1] = ""
			slice = CleanSlice(slice)
		} else if i != 0 && slice[i] == "(up)" {
			number = 1
			for j := 1; j <= number; j++ {
				if i-number >= 0 {
					if !IsWord(slice[i-j]) {
						number++
					}
					slice[i-j] = strings.ToUpper(slice[i-j])
				} else {
					break
				}
			}
			slice[i] = ""
			slice = CleanSlice(slice)
			i--
		} else if i != 0 && slice[i] == "(low)" {
			number = 1
			for j := 1; j <= number; j++ {
				if i-number >= 0 {
					if !IsWord(slice[i-j]) {
						number++
					}
					slice[i-j] = strings.ToLower(slice[i-j])
				} else {
					break
				}
			}
			slice[i] = ""
			slice = CleanSlice(slice)
			i--
		} else if i != 0 && slice[i] == "(cap)" {
			number = 1
			for j := 1; j <= number; j++ {
				if i-number >= 0 {
					if !IsWord(slice[i-j]) {
						number++
					}
					slice[i-j] = Capitalize(slice[i-j])
				} else {
					break
				}
			}
			slice[i] = ""
			slice = CleanSlice(slice)
			i--
		} else if i != 0 && slice[i] == "(hex)" {
			num, err := strconv.ParseInt(slice[i-1], 16, 64)

			if err != nil {
				slice[i] = ""
				slice = CleanSlice(slice)
				i--

			} else {
				slice[i-1] = strconv.Itoa(int(num))
				slice[i] = ""
				slice = CleanSlice(slice)
				i--
			}
		} else if i != 0 && slice[i] == "(bin)" {
			num, err := strconv.ParseInt(slice[i-1], 2, 64)

			if err != nil {
				slice[i] = ""
				slice = CleanSlice(slice)
				i--

			} else {
				slice[i-1] = strconv.Itoa(int(num))
				slice[i] = ""
				slice = CleanSlice(slice)
				i--
			}

		} else if i != 0 && i < len(slice)-1 && len(slice[i+1]) > 1 && slice[i] == "(up," && slice[i+1][len(slice[i+1])-1] == ')' && slice[i+1][len(slice[i+1])-2] != ')' {

			number, err := strconv.Atoi(slice[i+1][:len(slice[i+1])-1])

			if err != nil {
				fmt.Println("Input Invalid : ", err)
			} else {
				if number < 0 {
					number = 0
				}
				for j := 1; j <= number; j++ {
					if i-j >= 0 {
						if !IsWord(slice[i-j]) {
							number++
						}
						slice[i-j] = strings.ToUpper(slice[i-j])
					} else {
						break
					}
				}
				slice[i] = ""
				slice[i+1] = ""
				slice = CleanSlice(slice)
				i--
			}

		} else if i != 0 && i < len(slice)-1 && len(slice[i+1]) > 1 && slice[i] == "(low," && slice[i+1][len(slice[i+1])-1] == ')' && slice[i+1][len(slice[i+1])-2] != ')' {

			number, err := strconv.Atoi(slice[i+1][:len(slice[i+1])-1])

			if err != nil {
				fmt.Println("Input Invalid : ", err)
			} else {
				if number < 0 {
					number = 0
				}
				for j := 1; j <= number; j++ {
					if i-j >= 0 {
						if !IsWord(slice[i-j]) {
							number++
						}
						slice[i-j] = strings.ToLower(slice[i-j])
					} else {
						break
					}
				}
				slice[i] = ""
				slice[i+1] = ""
				slice = CleanSlice(slice)
				i--

			}

		} else if i != 0 && i < len(slice)-1 && len(slice[i+1]) > 1 && slice[i] == "(cap," && slice[i+1][len(slice[i+1])-1] == ')' && slice[i+1][len(slice[i+1])-2] != ')' {

			if CheckNumber(slice[i+1][:len(slice[i+1])-1]) {
				number, _ := strconv.Atoi(slice[i+1][:len(slice[i+1])-1])

			if number >= i {
				number = i
			}
				if number < 0 {
					number = 0
				}
				for j := 1; j <= number; j++ {
					if i-j >= 0 {
						if !IsWord(slice[i-j]) {
							number++
						}
						slice[i-j] = Capitalize(slice[i-j])
					} else {
						break
					}
				}
				slice[i] = ""
				slice[i+1] = ""
				slice = CleanSlice(slice)
				i--
			}
		}
	}
	 

	for i := 0; i < len(slice); i++ {
		if i != len(slice)-1 && (slice[i] == "a" || slice[i] == "A" || slice[i] == "'a" || slice[i] == "'A") {
			if IsVowels(slice[i+1]) {
				slice[i] += "n"
			}
		}
	}

	return slice
}
