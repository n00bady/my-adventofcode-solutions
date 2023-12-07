package main

import (
	"fmt"
	"regexp"
	"strconv"

	aoc "gitlab.com/n00bady/myaochelper"
)

func main() {
	fmt.Println("-= AoC 2023 Day 6 Part 1 =-")
	part1()
	fmt.Println()
	fmt.Println("-= AoC 2023 Day 6 Part 2 =-")
	part2()
}

func part2() {
	fmt.Println("Not implemented yet.")
}

func part1() {
	f := aoc.Read_Text_Input("./input.txt")
	input := aoc.Parse_Lines(f)

	var times []int
	var records []int

    // extract the numbers for the times and the records in order
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	for _, m := range re.FindAllString(input[0], -1) {
		tmp, _ := strconv.Atoi(m)
		times = append(times, tmp)
	}
	for _, m := range re.FindAllString(input[1], -1) {
		tmp, _ := strconv.Atoi(m)
		records = append(records, tmp)
	}

	fmt.Println("Time: ", times)
	fmt.Println("Distance: ", records)

    ans := 1 // 1 because we multiply

    // for each race 
	for r, t := range times {
		distance := 0
		win_count := 0
        // if holding it for 1ms or for tms then you don't move so I start from 1 and finish before t
        // calculate the distance for each try and if it's higher than the record increase the win_count
		for i := 1; i < t; i++ {
			distance = i * (t - i)
			if distance > records[r] {
				win_count++
			}
		}
        ans *= win_count // calculate the answer 
	}
    fmt.Println("Answer: ", ans)
}
