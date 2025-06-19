package goreloaded

import "fmt"

func IsPunctuatuion(s string) bool {
	result := true

	for i := 0; i < len(s); i++ {
		if s[i] == ',' || s[i] == '.' || s[i] == ';' || s[i] == ':'|| s[i] == '!' || s[i] == '?' {
			continue
		} else {
			result = false
		}
	}
	fmt.Println(result)
	return result
}
