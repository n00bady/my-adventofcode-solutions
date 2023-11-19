package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type coords struct {
	x int
	y int
}

func main() {
	fmt.Println("-=: AoC 2015 - Day 03 Part 1 :=-")
	var chars []rune

	f, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("File reading error: ", err)
		os.Exit(1)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		char, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Rune reading error: ", err)
			os.Exit(1)
		}

		chars = append(chars, char)
	}

	result := find_visited(chars)
	fmt.Println(result)
}

func find_visited(chars []rune) int {
	santa_loc := coords{
		x: 0,
		y: 0,
	}

	visited_coords := []coords{{x: 0, y: 0}}
	visited_count := 1

	for _, v := range chars {
		switch v {
		case '^':
			santa_loc.x++
		case 'v':
			santa_loc.x--
		case '<':
			santa_loc.y--
		case '>':
			santa_loc.y++
		}

		if !already_visited(santa_loc, visited_coords) {
			visited_coords = append(visited_coords, santa_loc)
			visited_count++
		}
	}

	return visited_count
}

func already_visited(santa_loc coords, visited_coords []coords) bool {
	for _, v := range visited_coords {
		if santa_loc == v {
			return true
		}
	}

	return false
}
