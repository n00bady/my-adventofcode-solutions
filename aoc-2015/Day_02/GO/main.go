package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("-=: AoC 2015 - Day 02 Part 1 :=-")

	// 2*l*w + 2*w*h + 2*h*l
	// l = length w = width h = heigth  In that order
	var total int

	input := "./input.txt"
	lines := parse_input(input)

	for _, line := range lines {
		lwh := strings.Split(line, "x")

		var dimensions [3]int
		dimensions[0], _ = strconv.Atoi(lwh[0]) // length
		dimensions[1], _ = strconv.Atoi(lwh[1]) // width
		dimensions[2], _ = strconv.Atoi(lwh[2]) // height

		smol_side := find_smol(dimensions)

		total += 2*dimensions[0]*dimensions[1] + 2*dimensions[1]*dimensions[2] + 2*dimensions[2]*dimensions[0] + smol_side
	}

	fmt.Println(total)
}

func find_smol(dimensions [3]int) int {
	var sides [3]int
	sides[0] = dimensions[0] * dimensions[1]
	sides[1] = dimensions[1] * dimensions[2]
	sides[2] = dimensions[2] * dimensions[0]
	smol := sides[0]
	for _, v := range sides {
		if v < smol {
			smol = v
		}
	}

	return smol
}

func parse_input(input string) []string {
	f, err := os.Open(input)
	if err != nil {
		fmt.Println("File reading error: ", err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
