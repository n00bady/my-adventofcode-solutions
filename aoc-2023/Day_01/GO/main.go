package main

import (
	"fmt"
	"strconv"
	"strings"
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
    // Although all my examples inputs seems to work 
    // I can't get correct output in the actual input
	f := aoc.Read_Text_Input("./input.txt")
	lines := aoc.Parse_Lines(f)
	sum := 0
	numbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for _, l := range lines {
		var digits [100]int
		var clean []int

		for n := range numbers {
			if strings.Index(l, n) > -1 {
				digits[strings.Index(l, n)] = numbers[n]
			}
		}
		for i, c := range l {
			if unicode.IsDigit(c) {
				digits[i] = int(c - '0')
			}
		}
		for _, v := range digits {
			if v != 0 {
				clean = append(clean, v)
			}
		}
		sum += clean[0]*10 + clean[len(clean)-1]

		fmt.Println(clean)
	}
	f.Close()

	fmt.Println("sum: ", sum)
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

	f.Close()

	fmt.Println("sum: ", sum)
}
