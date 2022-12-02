package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main()  {
  // Get input
  file, _ := os.Open("input.txt")
  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)
  var input []string
  for scanner.Scan() {
    input = append(input, scanner.Text())
  }
  file.Close()

  fmt.Println("-=Part 1=-")
  part1(input)
  fmt.Println("-=Part 2=-")
  part2(input)
}

func part1(input []string)  {
  var totals []int
  sum := 0
  num := 0

  // iterate each value and calculate their sum until 
  // there is a "" empty value then add the sum in a 
  // slice of ints 
  for _, c := range input {
    if c == "" {
      totals = append(totals, sum)
      sum = 0
    } else {
      num, _ = strconv.Atoi(c)
      sum += num
    }
  }
  // add the last sum at the end since there is no ""
  // empty value at the end
  totals = append(totals, sum)

  // Find and print the max value in the slice totals
  fmt.Println("max: ", slice_max(totals))
}

func part2(input []string) {
  // Same as part 1
  var totals []int
  sum := 0
  num := 0

  for _, c := range input {
    if c == "" {
      totals = append(totals, sum)
      sum = 0
    } else {
      num, _ = strconv.Atoi(c)
      sum += num
    }
  }
  totals = append(totals, sum)

  // sort all the sums
  sort.Ints(totals)
  length := len(totals)
  // then check the last 3 values and calculate their sum
  for i:=(length-1); i>=(length-3); i-- {
    sum += totals[i]
  }

  fmt.Println("Sum of top 3: ", sum)
}

// simple function to find max in a slice of ints
func slice_max(intslice []int) int {
  var max int

  for _, i := range intslice {
    if max < i {
      max = i
    }
  }
  
  return max
}
