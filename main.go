package main

import (
	"fmt"
	"os"
	"strings"

	"cohort/functions"
)

func main() {
	filename := os.Args

	if len(filename) != 3 {
		fmt.Println("invalid input")
		return
	}
	input := filename[1]
	output := filename[2]

	if (!strings.HasSuffix(input, ".txt") && !strings.HasSuffix(input, ".TXT")) || (!strings.HasSuffix(output, ".txt") && !strings.HasSuffix(output, ".TXT")) {
		fmt.Println("invalid input")
		return
	}

	text, err := os.ReadFile(filename[1])
	if err != nil {
		fmt.Println("error", err)
	}

	lines := strings.Split(string(text), "\n")
	var textfinal string

	for i := 0; i < len(lines); i++ {
		text2 := lines[i] + "\n"

		tret1 := goreloaded.SolvePunc(string(text2))
		tret2 := goreloaded.CleanStr(tret1)
		tret3 := goreloaded.StringToSlice(tret2)
		tret4 := goreloaded.CleanPunc(tret3)
		for goreloaded.ValidFlag(tret4) || goreloaded.ValidA(tret4) {
			tret4 = goreloaded.CleanPunc(goreloaded.ApplyFlag(tret4))
		}
		tretfinal := goreloaded.CleanCout(tret4, false)

		text3 := goreloaded.SliceToString(tretfinal)
		if i == len(lines)-1 {
			if len(text3) > 0 {
				textfinal += text3[:len(text3)-1]
			} else {
				textfinal += text3
			}
		} else {
			if len(text3) > 0 {
				textfinal += text3[:len(text3)-1] + "\n"
			} else {
				textfinal += text3 + "\n"
			}
		}

	}

	err = os.WriteFile(filename[2], []byte(textfinal), 0o644)
	if err != nil {
		fmt.Println("invalid input")
		return
	}
}
