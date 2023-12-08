package main

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	aoc "gitlab.com/n00bady/myaochelper"
)

// low              -                high
// 2, 3, 4, 5, 6, 7, 8, 9, T, J, Q, K, A

// Strongest To Weakest
// Five of a kind
// Four of a kind
// Full house (three of a kind + two of a kind)
// Three of a kind
// Two pair (2x two of a kind and the 5th doesn't matter)
// One pair (two of a kind)
// High card (everything is different)

// If 2 hands has the same type
// check the 1st card in each hand if they are different
// the hand with the highest card win
// if they are the same move to compare the seconds cards
// etc...

type card struct {
	symbol rune
	value  int
}

// h_type:
// 7 = five of a kind
// 6 = four of a kind
// 5 = full house
// 4 = three of a kind
// 3 = two pair
// 2 = one pair
// 1 = high card
type hand struct {
	cards   []card
	bid     int
	c_count map[card]int
	h_type  int
	rank    int
}

func main() {
	fmt.Println("-= AoC 2023 Day 7 Part 1 =-")
	part1()
	fmt.Println()
	fmt.Println("-= AoC 2023 Day 7 Part 2 =-")
	part2()
}

func part2() {
	fmt.Println("Not implemented yet.")
}

func part1() {
	f := aoc.Read_Text_Input("./input.txt")
	lines := aoc.Parse_Lines(f)

	// create a slice of hands from the input
	var hands []hand
	for _, l := range lines {
		str_fields := strings.Fields(l)
		var cards []card
		for _, c := range str_fields[0] {
			// create each card and add value to them depending on symbol
			if c >= '2' && c <= '9' {
				cards = append(cards, card{symbol: c, value: int(c - '0')})
			} else {
				switch c {
				case 'A':
					cards = append(cards, card{symbol: c, value: 14})
				case 'K':
					cards = append(cards, card{symbol: c, value: 13})
				case 'Q':
					cards = append(cards, card{symbol: c, value: 12})
				case 'J':
					cards = append(cards, card{symbol: c, value: 11})
				case 'T':
					cards = append(cards, card{symbol: c, value: 10})
				default:
					fmt.Println("Something is wrong I can feel it.")
				}
			}
		}
		bid, _ := strconv.Atoi(str_fields[1])

		hands = append(hands, hand{cards: cards, bid: bid, c_count: make(map[card]int), h_type: 1})
	}

	// Count the cards on each hand
	for _, h := range hands {
		for _, c := range h.cards {
			if slices.Contains(h.cards, c) {
				h.c_count[c]++
			}
		}
	}
	// This feels bad...
	// determine the type of each hand
	for i := 0; i < len(hands); i++ {
		pairs := 0
        has_ThreeOfaKind := false
		for _, v := range hands[i].c_count {
			switch v {
			case 5:
				hands[i].h_type = 7
				break
			case 4:
				hands[i].h_type = 6
				break
			case 3:
				hands[i].h_type = 4
                has_ThreeOfaKind = true
			case 2:
				hands[i].h_type = 2
				pairs++
			}
		}
		// two pair
		if pairs == 2 {
			hands[i].h_type = 3
		}
		// full house ???
		if pairs == 1 && has_ThreeOfaKind {
			hands[i].h_type = 5
		}
	}

	// sort the slice of hands
	sort.Slice(hands, func(i int, j int) bool {
		if hands[i].h_type == hands[j].h_type {
			for c := 0; c < 5; c++ {
				if hands[i].cards[c].value == hands[j].cards[c].value {
					continue
				}
				return hands[i].cards[c].value > hands[j].cards[c].value
			}
		}
		return hands[i].h_type > hands[j].h_type
	})

	for i := 0; i < len(hands); i++ {
		hands[i].rank = len(hands) - i
	}

	sum := 0
	for _, h := range hands {
		sum += (h.bid * h.rank)
	}
	fmt.Println("Sum: ", sum)
}
