package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
  var input []string
  file, _ := os.Open("input.txt")
  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)
  for scanner.Scan() {
    input = append(input, scanner.Text())
  }
  file.Close()
  fmt.Println(input)

  fmt.Println("-=Part 1=-")
  part1(input)
  fmt.Println("-=Part 2=-")
  part2(input)
  
}


func part1(input []string)  {
  // Iterate over the input and...
  length := len(input)
  counter := 0
  for i:=0; i<length; i++ {
    // splt each string on - and , then convert them to ints
    test := strings.FieldsFunc(input[i], delimiters)
    min1, _ := strconv.ParseInt(test[0], 10, 32)
    max1, _ := strconv.ParseInt(test[1], 10, 32)
    min2, _ := strconv.ParseInt(test[2], 10, 32)
    max2, _ := strconv.ParseInt(test[3], 10, 32)
    // check if they fully overlap and increase counter
    if min1 >= min2 && max1 <= max2 {
      counter++
    } else if min2 >= min1 && max2 <= max1 {
      counter++
    }
  }
  fmt.Println("Full Overlap Count: ", counter)

  
}

func part2(input []string) {

}

// helps me split a string on multiple delimiters 
// using string.FieldsFunc() method
func delimiters(r rune) bool {
  return r == '-' || r == ','
  
}
