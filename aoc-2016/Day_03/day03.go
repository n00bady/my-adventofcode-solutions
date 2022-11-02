package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
  var input [][]string
  file, _ := os.Open("input.txt")
  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)
  for scanner.Scan() {
    temp := make([]string, 0)
    temp = append(temp, scanner.Text())
    input = append(input, [][]string{temp}...)
  }
  file.Close()
  fmt.Println("-=Part_1=-")
  // run part1 func
  part1(input)
  fmt.Println("-=Part_2=-")
  // run part2 func
}

func checkifvalid(sides []int) bool {
  if sides[0] + sides[1] > sides[2] && sides[1] + sides[2] > sides[0] && sides[2] + sides[0] > sides[1] {
    return true
  }
  
  return false
}

func part1(input [][]string)  { 
  var t []int
  counter := 0

  for _, slice := range input {
    for _, line := range slice {
      s := strings.Fields(line)
      for _, n := range s {
        temp, _ := strconv.Atoi(n)
        t = append(t, temp)
      }
      // if t[0] + t[1] > t[2] && t[1] + t[2] > t[0] && t[2] + t[0] > t[1]{
      //   fmt.Println(t, " is correct!")
      //   counter += 1
      // }
      if checkifvalid(t) {
        counter += 1
      }
      t = nil
    }
  }

  fmt.Println(counter)
}
