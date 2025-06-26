package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	goreloaded "cohort"
)

func main() {
	filename := os.Args

	if len(filename) != 3 || filename[2][len(filename[2])-4:len(filename[2])] != ".txt" {
		fmt.Println("invalid input")
		return
	}

	text, err := os.ReadFile(filename[1])
	if err != nil {
		fmt.Println("error", err)
	}

	lines := strings.Split(string(text), "\n")

	textf := ""

	for _, line := range lines {
		text2 := line + "\n"

		text3 := goreloaded.CleanStr(SolvePunc(string(text2)))
		fmt.Println(text3)
		text4 := CleanPunc(goreloaded.StringToSlice(text3), false)
		for goreloaded.ValidFlag(text4) {
				text4 = GoReloaded(text4)
				fmt.Println(text4)
		}
		
		text5 := goreloaded.SliceToString(text4)
		textf += text5 + "\n"
	}

	err = os.WriteFile(filename[2], []byte(textf), 0o644)
	if err != nil {
		fmt.Println("invalid input")
		return
	}
}

func CleanPunc(slice []string, flag bool) []string {
	count := 0
	for i := 0; i < len(slice); i++ {

		if len(slice[i]) != 0 && slice[i] != "'" && flag {

			if slice[i][0] == '\'' {
				count++
			}
			if slice[i][len(slice[i])-1] == '\'' {
				count++
			}
		}
		if i != len(slice)-1 && slice[i] == "'" && count%2 == 0 && slice[i+1][0] != '(' {
			if slice[i+1] == "'" {
				count++
			}
			slice[i+1] = slice[i] + slice[i+1]
			slice[i] = ""
			slice = goreloaded.CleanSlice(slice)
			i--
			count++
		} else if i != 0 && slice[i] == "'" && count%2 != 0 && slice[i-1][len(slice[i-1])-1] != ')' {
			if slice[i-1] == "'" {
				count++
			}
			slice[i-1] += slice[i]
			slice[i] = ""
			slice = goreloaded.CleanSlice(slice)
			i--
			count++
		} else if i != 0 && goreloaded.IsPunctuation(rune(slice[i][0])) && slice[i-1][len(slice[i-1])-1] != ')' {
			slice[i-1] += slice[i]
			slice[i] = ""
			slice = goreloaded.CleanSlice(slice)
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
			 if i != 0 && !goreloaded.IsPunctuation(s[i-1]) {
				str += " " + string(s[i])
			} else if i != len(s)-1 && !goreloaded.IsPunctuation(s[i+1]) {
				str += string(s[i]) + " "
			} else {
				str += string(s[i])
			}
		} else if i != 0 && s[i] == '(' && goreloaded.IsPunctuation(s[i-1]) {
			str += " " + string(s[i])
		} else if i != len(s)-1 && s[i] == ')' && goreloaded.IsPunctuation(s[i+1]) {
			str += string(s[i]) + " "
		} else if i != 0 && i != len(s)-1 && (s[i] == '\'' && ((goreloaded.IsPunctuation(s[i-1]) || goreloaded.IsPunctuation(s[i+1]) || s[i-1] == ' ' ||
			s[i+1] == ' ') || s[i+1] == '\'' || s[i-1] == '\'')) {
			str += " " + string(s[i]) + " "
		} else if s[i] == '\'' && (i == 0 || i == len(s)-1) {
			str += " " + string(s[i]) + " "
		} else {
			str += string(s[i])
		}
	}

	return str
}

func GoReloaded(slice []string) []string {
	for i := 0; i < len(slice); i++ {
		if slice[0] == "(up)" || slice[0] == "(low)" || slice[0] == "(cap)" || slice[0] == "(hex)" || slice[0] == "(bin)" {
			slice[0] = ""
			slice = goreloaded.CleanSlice(slice)
			i--
		} else if len(slice) > 1 && (slice[0] == "(up," || slice[0] == "(low," || slice[0] == "(cap,") && goreloaded.CheckNumber(slice[1][:len(slice[1])-1]) && slice[1][len(slice[1])-1] == ')' {
			fmt.Println(goreloaded.CheckNumber(slice[1][:len(slice[1])-1]))
			fmt.Println(slice[1][:len(slice[1])-1])
			slice[0] = ""
			slice[1] = ""
			slice = goreloaded.CleanSlice(slice)
			i --
		} else if i != 0 && slice[i] == "(up)" {
			slice[i-1] = strings.ToUpper(slice[i-1])
			slice[i] = ""
			slice = goreloaded.CleanSlice(slice)
			i--
		} else if i != 0 && slice[i] == "(low)" {
			slice[i-1] = strings.ToLower(slice[i-1])
			slice[i] = ""
			slice = goreloaded.CleanSlice(slice)
			i--
		} else if i != 0 && slice[i] == "(cap)" {
			slice[i-1] = goreloaded.Capitalize(slice[i-1])
			slice[i] = ""
			slice = goreloaded.CleanSlice(slice)
			i--
		} else if i != 0 && slice[i] == "(hex)" {
			num, err := strconv.ParseInt(slice[i-1], 16, 64)

			if err != nil {
				fmt.Printf("Error converting '%s': %v\n", slice[i-1], err)
				slice[i] = ""
				slice = goreloaded.CleanSlice(slice)
				i--
			} else {
				slice[i-1] = strconv.Itoa(int(num))
				slice[i] = ""
				slice = goreloaded.CleanSlice(slice)
				i--
			}
		} else if i != 0 && slice[i] == "(bin)" {
			num, err := strconv.ParseInt(slice[i-1], 2, 64)

			if err != nil {
				fmt.Printf("Error converting '%s': %v\n", slice[i-1], err)
				slice[i] = ""
				slice = goreloaded.CleanSlice(slice)
				i--
			} else {
				slice[i-1] = strconv.Itoa(int(num))
				slice[i] = ""
				slice = goreloaded.CleanSlice(slice)
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
						slice[i-j] = strings.ToUpper(slice[i-j])
					} else {
						break
					}
				}
				slice[i] = ""
				slice[i+1] = ""
				slice = goreloaded.CleanSlice(slice)
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
						slice[i-j] = strings.ToLower(slice[i-j])
					} else {
						break
					}
				}
				slice[i] = ""
				slice[i+1] = ""
				slice = goreloaded.CleanSlice(slice)
				i--
			}

		} else if i != 0 && i < len(slice)-1 && len(slice[i+1]) > 1 && slice[i] == "(cap," && slice[i+1][len(slice[i+1])-1] == ')' && slice[i+1][len(slice[i+1])-2] != ')' {

			number, err := strconv.Atoi(slice[i+1][:len(slice[i+1])-1])

			if err != nil {
				fmt.Println("Input Invalid : ", err)
			} else {
				if number < 0 {
					number = 0
				}
				for j := 1; j <= number; j++ {
					if i-j >= 0 {
						slice[i-j] = goreloaded.Capitalize(slice[i-j])
					} else {
						break
					}
				}
				slice[i] = ""
				slice[i+1] = ""
				slice = goreloaded.CleanSlice(slice)
				i--
			}

		}
	}

	for i := 0; i < len(slice); i++ {
		if i != len(slice)-1 && (slice[i] == "a" || slice[i] == "A" || slice[i] == "'a" || slice[i] == "'A") {
			if goreloaded.IsVowels(slice[i+1]) {
				slice[i] += "n"
			}
		}
	}
	slice = CleanPunc(slice, true)
	return slice
}
