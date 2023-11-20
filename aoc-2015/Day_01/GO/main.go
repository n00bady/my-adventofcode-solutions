package main

import (
	"fmt"

	aoc "gitlab.com/n00bady/myaochelper"
)

func main() {
	fmt.Println("-=: AoC 2015 - Day 01 Part 1 :=-")

	f := aoc.Read_Text_Input("./input.txt")
	chars := aoc.Parse_Runes(f)
    f.Close()
    floor := checkChar(chars)

	fmt.Println("You end up in floor: ", floor)
}

func checkChar(chars []rune) int {
	current_floor := 0

	for _, c := range chars {
		if c == '(' {
			current_floor++
		} else if c == ')' {
			current_floor--
		}
	}

	return current_floor
}
