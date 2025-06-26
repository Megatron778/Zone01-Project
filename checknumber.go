package goreloaded

import ("fmt" 
 "strconv")

func CheckNumber(nb string) bool {
	_, err := strconv.Atoi(nb)

			if err != nil {
				fmt.Println("Input Invalid : ", err)
				return false
			} else {
				return true
	}
}
