package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	aoc "gitlab.com/n00bady/myaochelper"
)

func main() {
	fmt.Println("-= AoC 2023 Day 5 Part 1 =-")
	part1()
	fmt.Println()
	fmt.Println("-= AoC 2023 Day 5 Part 2 =-")
	part2()
}

func part2() {
	fmt.Println("Not implemented yet.")
}

// This is not performant at all when running on the actual input
// my cpu and ram maxes and then it crashes :(
func part1() {
	f := aoc.Read_Text_Input("./input.txt")
	lines := aoc.Parse_Lines(f)

	var seeds []int

	// regex to extract numbers
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	// get the seeds first
	for _, m := range re.FindAllString(lines[0], -1) {
		tmp, _ := strconv.Atoi(m)
		seeds = append(seeds, tmp)
	}
	// a copy of the seeds slice (I dunno I might need the seeds later...)
	converted := make([]int, len(seeds))
	copy(converted, seeds)

	idx := 0
	idx_next := 0
	sections := []string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"}
	for i, s := range sections {
		// find the line for each map
		idx = FindIndex(lines, s)
		if i >= 6 {
			idx_next = len(lines)
		} else {
			idx_next = FindIndex(lines, sections[i+1])
		}
		// for every section check each seed untill you find it in the number entries or it's the end
		for c, s := range converted {
		next_seed:
			for j := idx + 1; j < idx_next; j++ {
				// Extract the numbers between each map
				var nums []int
				for _, m := range re.FindAllString(lines[j], -1) {
					tmp, _ := strconv.Atoi(m)
					nums = append(nums, tmp)
				}
				// in case there are extra lines without numbers
				if len(nums) > 0 {
					// define the range
					dest_start := nums[0]
					src_start := nums[1]
					rng := nums[2]
					// and convert every seed
					if s >= src_start && s <= src_start+rng-1 {
						// That's the formula... took me way too long...
						// new seed = (seed - start) + dest
						converted[c] = (converted[c] - src_start) + dest_start
						break next_seed
					}
				}
			}
		}
        // oof 4 nested fors :S
	}

	fmt.Println("The seeds: ", seeds)
	fmt.Println()
	fmt.Println("Converted to locations: ", converted)
	fmt.Println()
	fmt.Println("Smallest location: ", converted[FindMinIndex(converted)])
}

// This could go into myaochelper
func FindIndex(lines []string, key string) int {
	for i, l := range lines {
		if strings.Contains(l, key) {
			return i
		}
	}

	return -1
}

// THis could go in myaochelper
func FindMinIndex(s []int) int {
	min := s[0]
	minIdx := 0
	for i, v := range s {
		if v < min {
			min = v
			minIdx = i
		}
	}

	return minIdx
}
