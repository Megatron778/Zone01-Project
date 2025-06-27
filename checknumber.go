package goreloaded

import (
 "strconv")

func CheckNumber(nb string) bool {
	_, err := strconv.Atoi(nb)

			if err == nil {
				return true
			} else {
				return false
	}
}
