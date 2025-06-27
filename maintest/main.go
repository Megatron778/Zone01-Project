package main

import (
	"fmt"
	"os"
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

		text3 := goreloaded.CleanStr(goreloaded.SolvePunc(string(text2)))
		text4 := goreloaded.CleanPunc(goreloaded.StringToSlice(text3))
		for goreloaded.ValidFlag(text4) || goreloaded.ValidA(text4) {
			text4 = goreloaded.CleanPunc(goreloaded.GoReloaded(text4))
		}
		texta := goreloaded.CleanCout(text4, false)

		text5 := goreloaded.SliceToString(texta)
		textf += text5 + "\n"
	}

	err = os.WriteFile(filename[2], []byte(textf), 0o644)
	if err != nil {
		fmt.Println("invalid input")
		return
	}
}
