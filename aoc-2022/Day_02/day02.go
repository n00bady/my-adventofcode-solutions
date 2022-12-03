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
  scanner.Split(bufio.ScanLines)
  for scanner.Scan() {
    input = append(input, scanner.Text())
  }
  file.Close()

  fmt.Println("-=Part 1=-") 
  part1(input)
  fmt.Println("-=Part 2=-")
  part2()
}

func part1(input []string)  {
  var op []int
  var me []int
  score := 0

  // create 2 slices one for the opponent and one for me and substitute
  // the letters with numbers(I guess you don't need to do that but it's easier for me to think like that?)
  for _, c := range input {
    switch c[0] {
      case 'A': op = append(op, 1)
      case 'B': op = append(op, 2)
      case 'C': op = append(op, 3)
    }
    switch c[2] {
      case 'X': me = append(me, 1)
      case 'Y': me = append(me, 2)
      case 'Z': me = append(me, 3)  
    }
  }

  // find the length for one of the slices I created above 
  // I could check if they have the same length but I assume
  // the input is always correct so I ingore the errors
  length_op := len(op)
  // iterate of the slices, check who wins and calculate score
  for i:=0; i<length_op; i++ {
    if op[i] == me[i] {
      score += me[i] + 3
      // This could be done better but whatever...
    } else if (op[i] == 3 && me[i] == 1) || (op[i] == 1 && me[i] == 2) || (op[i] == 2 && me[i] == 3) {
      score += me[i] + 6
    } else {
      score += me[i]
    }
  }
  fmt.Println("Final Score: ", score)

}

func part2()  {
  fmt.Println("Not yet!")
}

