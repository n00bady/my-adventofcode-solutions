package main

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode"

	aoc "gitlab.com/n00bady/myaochelper"
)

func main() {
	fmt.Println("-= AoC 2023 Day 01 Part 1")
	part1()
	fmt.Println("-= AoC 2023 Day 01 Part 2")
    part2()
}

func part2() {
    // I have no idea how to make it match eighthree as eight and three ???
    m := regexp.MustCompile(`((f(ive|our)|s(even|ix)|t(hree|wo)|(ni|o)ne|eight))`)
	f := aoc.Read_Text_Input("./test_input_2.txt")
	lines := aoc.Parse_Lines(f)

    for _, l := range lines {
        res := m.FindAllString(l, 10)
        fmt.Println(res)
    }
}

func part1() {
	f := aoc.Read_Text_Input("./input.txt")
	lines := aoc.Parse_Lines(f)
	sum := 0

	for _, l := range lines {
		var digits []rune
		for _, c := range l {
			if unicode.IsDigit(c) {
				digits = append(digits, c)
			}
		}
		num, err := strconv.Atoi(fmt.Sprintf("%c%c", digits[0], digits[len(digits)-1]))
		if err != nil {
			fmt.Println(err)
		}
		sum += num
	}

	fmt.Println("sum: ", sum)
}
