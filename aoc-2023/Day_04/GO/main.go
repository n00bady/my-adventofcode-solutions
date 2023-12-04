package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	aoc "gitlab.com/n00bady/myaochelper"
)

func main() {
	fmt.Println("-= AoC 2023 Day 4 Part 1 =-")
	part1()
	fmt.Println()
	fmt.Println("-= AoC 2023 Day 4 Part 2 =-")
	part2()
}

func part2() {
	fmt.Println("Not implemented yet.")
}

func part1() {
	f := aoc.Read_Text_Input("./input.txt")
	lines := aoc.Parse_Lines(f)

	// split each line in Game ##:
	var cards []string
	for _, l := range lines {
		cards = append(cards, strings.Split(l, ":")[1])
	}

	var win_strs []string
	var lucky_strs []string
	var win_nums [][]int
	var lucky_nums [][]int
	for _, c := range cards {
		win_lose_split := strings.Split(c, "|")
		win_strs = append(win_strs, win_lose_split[0])
		lucky_strs = append(lucky_strs, win_lose_split[1])
	}

	// when it a single digit number it breaks my slices because it return 0 for the space character!!!
	// why they couldn't write 06 instead of <whitespace>6 ?!?!?! solved it by just appending only when
    // err == nil 

	// These are for the left side of | then winning numbers
	for _, s := range win_strs {
		tmp := strings.Split(strings.TrimSpace(s), " ")
		var tmp_nums []int
		for _, n := range tmp {
			number, err := strconv.Atoi(n)
			if err == nil {
				tmp_nums = append(tmp_nums, number)
			}
			if err != nil {
				fmt.Println("Cannot convert ", n, " to int! ", err)
			}
		}
		win_nums = append(win_nums, tmp_nums)
	}

	// These are for the right side of | the "lucky" numbers
	for _, s := range lucky_strs {
		tmp := strings.Split(strings.TrimSpace(s), " ")
		var tmp_nums []int
		for _, n := range tmp {
			number, err := strconv.Atoi(n)
			if err == nil {
				tmp_nums = append(tmp_nums, number)
			}
			if err != nil {
				fmt.Println("Cannot convert ", n, " to int!", err)
			}
		}
		lucky_nums = append(lucky_nums, tmp_nums)
	}

    // since every slice is in order we just check to see
    // how many win_nums are in lucky_nums
    // and we sum the 2^count-1 to find the result
    sum := 0
	for i := 0; i < len(win_nums); i++ {
		count := 0
		for j := 0; j < len(win_nums[i]); j++ {
			if slices.Contains(lucky_nums[i], win_nums[i][j]) {
				count++
			}
		}
        sum += int(math.Pow(2, float64(count-1))) // if count-1 = -1 then the int() makes it 0 ?
	}
    fmt.Println(sum)
}
