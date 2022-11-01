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
        case 'U': if y > 0 { y -= 1 }
        case 'D': if y < 2 { y += 1 }
        case 'L': if x > 0 { x -= 1 }
        case 'R': if x < 2 { x += 1 }
      }
    }
    finalpos = append(finalpos, keypad[y][x])
  }
  fmt.Println("-=Part_1=-")
  fmt.Println(finalpos)
  fmt.Println("-=Part_2=-")
  part2()
}


func part2()  {
  keypad := [][]int {{0, 0, 1, 0, 0}, {0, 2, 3, 4, 0}, {5, 6, 7, 8, 9}, {0, 0xA, 0xB, 0xC, 0}, {0, 0, 0xD, 0, 0}}
  
  file, _ := os.Open("input.txt")
  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)
  var input []string
  for scanner.Scan() {
    input = append(input, scanner.Text())
  }
  file.Close()
  
  x, y := 2, 0
  var finalpos []int

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
