package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"unicode"

	aoc "gitlab.com/n00bady/myaochelper"
)

type coord struct {
	x int
	y int
}

var visited []coord

func main() {
	fmt.Println("-= AoC 2023 Day 02 Part 1 =-")
	part1()
	fmt.Println()
	fmt.Println("-= AoC 2023 Day 02 Part 2 =-")
	part2()
}

func part2() {
	fmt.Println("Not implemented yet.")
}

func part1() {
	f := aoc.Read_Text_Input("./input.txt")
	lines := aoc.Parse_Lines(f)
	f.Seek(0, io.SeekStart)
	lc := LineCounter(f)
	input := make([][]rune, lc)
	sum := 0
	// I though it was going to help
	// but I think I just make more complex that it needs to be :(

	coords := make(map[coord]rune)
	for i, l := range lines {
		for j, r := range l {
			input[i] = append(input[i], r)
			coords[coord{x: i, y: j}] = r
		}
	}

	adjPoints := []coord{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}
	// There has to be a better way...
	for k, v := range coords {
		// if it's not a number or . then it's a symbol
		// search the adjacnet positions for digits
		if !unicode.IsDigit(v) && v != '.' {
			// check adjucent positions
			/*
			   {-1, -1} {-1, 0} {-1, 1}
			   {0,  -1} {  P  } {0,  1}
			   {1,  -1} {1,  0} {1,  1}
			*/
			for _, p := range adjPoints {
				// if it's a digit and not in the visited coords then it has to be what we want
				if unicode.IsDigit(coords[coord{k.x + p.x, k.y + p.y}]) && !slices.Contains(visited, coord{k.x + p.x, k.y + p.y}) {
					visited = append(visited, coord{k.x + p.x, k.y + p.y})
					num := makeNumb(coords, coord{k.x + p.x, k.y + p.y})
					// we played enough with runes now it's time for actual numbers
					actual_number, err := strconv.Atoi(string(num))
					if err != nil {
						fmt.Println("That is not a number! ", err)
					}
					sum += actual_number
				}
			}
		}
	}
	// fuck this years day 3!

	fmt.Println("Sum: ", sum)
}

// Take a starting coord and check left and right for other digits to make a number
func makeNumb(coords map[coord]rune, startPos coord) []rune {
	var num []rune

	// to the left
	x := startPos.x
	y := startPos.y
	for unicode.IsDigit(coords[coord{x, y}]) {
		num = append([]rune{coords[coord{x, y}]}, num...)
		visited = append(visited, coord{x, y})
		y--
	}

	// to the right
	x = startPos.x
	y = startPos.y + 1
	for unicode.IsDigit(coords[coord{x, y}]) {
		num = append(num, coords[coord{x, y}])
		visited = append(visited, coord{x, y})
		y++
	}

	return num
}

// This probably should be included in myaochelper
func LineCounter(f *os.File) int {
	count := 0
	buf := make([]byte, bufio.MaxScanTokenSize)

	for {
		bufferSize, err := f.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return 0
		}

		var buff_pos int
		for {
			i := bytes.IndexByte(buf[buff_pos:], '\n')
			if i == -1 || bufferSize == buff_pos {
				break
			}
			buff_pos += i + 1
			count++
		}
		if err == io.EOF {
			break
		}
	}

	return count
}
