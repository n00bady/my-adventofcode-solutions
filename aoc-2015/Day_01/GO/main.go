package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("-=: AoC 2015 - Day 01 Part 1 :=-")
    fmt.Println()

	var current_floor int

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

		current_floor = checkChar(char, current_floor)
	}

	fmt.Println("You end up in floor: ", current_floor)
}

func checkChar(char rune, current_floor int) int {

	if char == '(' {
        current_floor++
	} else if char == ')' {
        current_floor--
	}

	return current_floor
}
