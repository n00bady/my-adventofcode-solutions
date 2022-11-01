package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
  // Get input from the file
  file, _ := os.Open("input.txt")
  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)
  // Scan it line by line and put them in a string slice
  var input []string
  for scanner.Scan() {
    input = append(input, scanner.Text())
  }
  file.Close()

  fmt.Println("-=Part_1=-")
  part1(input)
  fmt.Println("-=Part_2=-")
  part2(input)
}

func part1(input []string) {
  // Constract a 2D array that represends with the an ordinary keypad
  // and use the keypad numbers as values
  keypad := [][]int {{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
  x, y := 1, 1
  var finalpos []int

  // iterate through the input line by line
  for _, line := range input {
    // and then character by character
    for _, ch := range line {
      // and follow the moves by increasing and decreasing the indexes
      // according to the movements 
      switch ch {
        case 'U': if x > 0 { x -= 1 }
        case 'D': if x < 2 { x += 1 }
        case 'L': if y > 0 { y -= 1 }
        case 'R': if y < 2 { y += 1 }
      }
    }
    // build the final solution digit by digit
    finalpos = append(finalpos, keypad[x][y])
  }

  fmt.Println(finalpos)
}

func part2(input []string)  {
  // Construct a 2d array 5x5 the posisitons that the new keypad doesn't have
  // a key are represended by 0s and are ignored
  keypad := [][]int {{0, 0, 1, 0, 0}, {0, 2, 3, 4, 0}, {5, 6, 7, 8, 9}, {0, 0xA, 0xB, 0xC, 0}, {0, 0, 0xD, 0, 0}}
  x, y := 2, 0
  var finalpos []int

  // I do the same things as part one only with the additional condition 
  // that if the the value is 0 at that position go back again 😝
  for _, line := range input {
    for _, ch := range line {
      switch ch {
      case 'U': if x > 0 { x -= 1; if keypad[x][y] == 0 { x += 1 } }
      case 'D': if x < 4 { x += 1; if keypad[x][y] == 0 { x -= 1 } }
      case 'L': if y > 0 { y -= 1; if keypad[x][y] == 0 { y += 1 } }
      case 'R': if y < 4 { y += 1; if keypad[x][y] == 0 { y -= 1 } }
      }
    }
    finalpos = append(finalpos, keypad[x][y])
  }
  fmt.Printf("%x\n", finalpos)
}
