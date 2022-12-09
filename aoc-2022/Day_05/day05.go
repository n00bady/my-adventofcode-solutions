package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []string

func main()  {
  var input []string
  file, _ := os.Open("input.txt")
  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)
  for scanner.Scan() {
    input = append(input, scanner.Text())
  }
  file.Close()

  // initialize the stack
  p1 := StartingStackInit()
  fmt.Println("-=Part 1=-")
  part1(input, p1)
  fmt.Println("-=Part 2=-")
  part2(input)
  
  
}

func part1(input []string, s [9]Stack) {
  // Iterate of each line (I have cut out the first part of the input thas has the initial state of the stacks)
  // and parse the numbers and convert them to integers
  for _, line := range input {
    l := strings.Split(line, " ")
    move, _ := strconv.ParseInt(l[1], 10, 0)
    orig, _ := strconv.ParseInt(l[3], 10, 0)
    dest, _ := strconv.ParseInt(l[5], 10, 0)
    // loop as many times as the move number
    for p:=0; p<int(move); p++ {
      // and the top of the stack to to tmp variable
      top := s[int(orig)-1].Top()
      // delete it and then add the top to the destination
      s[int(orig)-1].Pop()
      s[int(dest)-1].Push(top)
    }
  }
  // make a string slice and append all the top values for every stack and print them
  var rslt []string
  for i:=0; i<9; i++ {
    rslt = append(rslt, s[i].Top())
  }
  fmt.Println("All the tops: ", rslt)
}

func part2(input []string)  {
  
}

//     [V] [G]             [H]        
// [Z] [H] [Z]         [T] [S]        
// [P] [D] [F]         [B] [V] [Q]    
// [B] [M] [V] [N]     [F] [D] [N]    
// [Q] [Q] [D] [F]     [Z] [Z] [P] [M]
// [M] [Z] [R] [D] [Q] [V] [T] [F] [R]
// [D] [L] [H] [G] [F] [Q] [M] [G] [W]
// [N] [C] [Q] [H] [N] [D] [Q] [M] [B]
//  1   2   3   4   5   6   7   8   9 
// I use these functions to initialize the slice of stacks on their
// initial state
func StartingStackInit() [9]Stack {
  var s [9]Stack
  crates := [][]string{{"N", "D", "M", "Q", "B", "P", "Z"},
                       {"C", "L", "Z", "Q", "M", "D", "H", "V"},
                       {"Q", "H", "R", "D", "V", "F", "Z", "G"},
                       {"H", "G", "D", "F", "N"},
                       {"N", "F", "Q"},
                       {"D", "Q", "V", "Z", "F", "B", "T"},
                       {"Q", "M", "T", "Z", "D", "V", "S", "H"},
                       {"M", "G", "F", "P", "N", "Q"},
                       {"B", "W", "R", "M"}}
  for i, v := range crates {
    for _, c := range v {
      s[i].Push(c)
    }
  }
  
  return s
}

//     [D]    
// [N] [C]    
// [Z] [M] [P]
//  1   2   3
func TestStackInit() [3]Stack {
  var s [3]Stack

  s[0].Push("Z")
  s[0].Push("N")
  
  s[1].Push("M")
  s[1].Push("C")
  s[1].Push("D")

  s[2].Push("P")

  return s
}

// ---

// Stack sruct methods

func (s *Stack) Push(v string)  {
  *s = append(*s, v)
}

func (s *Stack) Pop() string {
  ret := (*s)[len(*s)-1]
  *s = (*s)[0:len(*s)-1]

  return ret
}

func (s *Stack) Top() string {

  return (*s)[len(*s)-1]
}
