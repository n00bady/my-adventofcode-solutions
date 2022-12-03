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
  part2(input)
 
}

func part1(input []string)  {
  sum := 0
  // take the input and make a slice of characters
  for _, r := range input {
    var tmpsack []rune
    for _, c := range r {
      tmpsack = append(tmpsack, c)
    }
    
    var compartmentA []rune
    var compartmentB []rune
    length := len(tmpsack)
    // split the slice of chars in half and save it in 2 different variables
    for i:=0; i<(length/2); i++ {
      compartmentA = append(compartmentA, tmpsack[i])
    }
    for i:=(length/2); i<length; i++ {
      compartmentB = append(compartmentB, tmpsack[i])
    }
    
    // iterate over all the elements of the slice A and check if it exist in slice B
    // if it is calculate the prioritiy and add it to the sum
    for i:=0; i<(length/2); i++ {
      if Contains(compartmentB, compartmentA[i]) {
        if compartmentA[i] >= 'a' && compartmentA[i] <= 'z' {
          sum += int(compartmentA[i]) - 96
        } else if compartmentA[i] >= 'A' && compartmentA[i] <= 'Z'{
          sum += int(compartmentA[i]) - 38
        }
        break
      }
    }
  }
  fmt.Println("Sum: ", sum)
}

func part2(input []string)  {
  fmt.Print("Not yet!")
  
}

// Helps in checking if a slice contains an element
func Contains(slice []rune, e rune) bool {
  for _, r := range slice {
    if r == e {
      return true
    }
  } 

  return false
}
