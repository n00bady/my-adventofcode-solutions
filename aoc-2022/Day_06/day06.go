package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
  var input []string
  file, _ := os.Open("input.txt")
  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanRunes)
  for scanner.Scan() {
    input = append(input, scanner.Text())
  }

  fmt.Println("-=Part 1=-")
  part1(input)
  fmt.Println("-=Part 2=-")
  part2(input)
}

func part1(input []string) int {
  // start to iterate from the 4th element
  for i:=3; i<len(input); i++ {
    // create a temporary slice of strings and
    // put insed the last 4 elements
    var tmp []string
    for j:=3; j>=0; j-- {
      tmp = append(tmp, input[i-j])
    }
    // check if the tmp slice's elements are unique
    // if they are print the index(+1 cause we start from 0) and exit
    if areDistinct(tmp) {
      fmt.Println(i+1)
      return 0
    }
  }

  return 1
}

func part2(input []string) int {
  // pretty much the same thing we did for part1
  // but now we check the last 14 elements
  for i:=13; i<len(input); i++ {
    var tmp []string
    for j:=13; j>=0; j-- {
      tmp = append(tmp, input[i-j])
    }
    if areDistinct(tmp) {
      fmt.Println(i+1)
      return 0
    }
  } 
  
  return 1
}

// helper function to check if all
// the elements of a slice are distinct
func areDistinct(s []string) bool {
  freq := make(map[string]int)
    for _, e := range s {
      freq[e] += 1
  }
  
  return len(s) == len(freq)
}

