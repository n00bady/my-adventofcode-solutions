package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
  file, _ := os.Open("input.txt")
  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)
  var input []string
  for scanner.Scan() {
    input = append(input, scanner.Text())
  }
  file.Close()

  keypad := [][]int {{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
  x, y := 1, 1
  var finalpos []int

  for _, lin := range input {
    for _, c := range lin {
      switch c {
        case 'D': if y < 2 { y += 1 }
        case 'U': if y > 0 { y -= 1 }
        case 'L': if x > 0 { x -= 1 }
        case 'R': if x < 2 { x += 1 }
      }
    }
    finalpos = append(finalpos, keypad[y][x])
  }
  fmt.Println("-=Part_1=-")
  fmt.Println(finalpos)
  fmt.Println("-=Part_2=-")
  fmt.Println("to be contineued...")
}
