package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "gitlab.com/n00bady/myaochelper"
)

func main() {
	fmt.Println("-= AoC 2023 Day 02 Part 1 =-")
	part1()
	fmt.Println()
	fmt.Println("-= AoC 2023 Day 02 Part 2 =-")
	part2()
}

func part2() {
	fmt.Println("Not implemented yet...")
}

func part1() {
	f := aoc.Read_Text_Input("./input.txt")
	input := aoc.Parse_Lines(f)
	sum := 0

	for id, l := range input {

		nogame := strings.Split(l, ":")          // ignore the Game ##:
		subsets := strings.Split(nogame[1], ";") // Split it in subsets
		check_flag := true

		for _, s := range subsets {
			rgb_counts := []int{0, 0, 0} // 1st Red then Green and last Blue
			color_count := strings.Split(s, ",")
			for _, c := range color_count {
				col_amounts := strings.Split(c, " ")
				switch col_amounts[2] {
				case "red":
					rgb_counts[0], _ = strconv.Atoi(col_amounts[1])
				case "green":
					rgb_counts[1], _ = strconv.Atoi(col_amounts[1])
				case "blue":
					rgb_counts[2], _ = strconv.Atoi(col_amounts[1])
				}
				if !check_limits(rgb_counts) {
					check_flag = false
					break
				}
			}
		}
		if !check_flag {
			check_flag = true
		} else {
			sum += id + 1
		}
	}

	fmt.Println("The sum of valid games: ", sum)
}

// returns True if the values are within the limits
func check_limits(rgb_counts []int) bool {
	red_limit := 12
	green_limit := 13
	blue_limit := 14

	if rgb_counts[0] > red_limit {
		return false
	}
	if rgb_counts[1] > green_limit {
		return false
	}
	if rgb_counts[2] > blue_limit {
		return false
	}

	return true
}
