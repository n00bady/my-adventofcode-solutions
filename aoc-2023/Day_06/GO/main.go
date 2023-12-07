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

// Doing it the same way I do it in the 1st part
// I guess there are a way to optimize it but on my modest Computer
// runs almost instantly and I am a mathlet so I doubt I can find
// the way to optimize it any time soon...
func part2() {
	f := aoc.Read_Text_Input("./input.txt")
	input := aoc.Parse_Lines(f)

	var time int
	var record int

	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	var str_time string
	for _, m := range re.FindAllString(input[0], -1) {
		str_time += m
	}
	time, _ = strconv.Atoi(str_time)
	var str_record string
	for _, m := range re.FindAllString(input[1], -1) {
		str_record += m
	}
	record, _ = strconv.Atoi(str_record)
	fmt.Println("Time: ", time)
	fmt.Println("Record: ", record)

	distance := 0
	win_count := 0
	for i := 1; i < time; i++ {
		distance = i * (time - i)
		if distance > record {
			win_count++
		}
	}
	fmt.Println("Wind count: ", win_count)
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
