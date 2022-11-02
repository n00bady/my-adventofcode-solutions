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
  part1(input)
  fmt.Println("-=Part_2=-")
  part2(input)
}

// check is a the 3 sides are a valid triangle 
// it takes as input a slice of 3 ints and returns true or false
func checkifvalidTriangle(sides []int) bool {
  if sides[0] + sides[1] > sides[2] && sides[1] + sides[2] > sides[0] && sides[2] + sides[0] > sides[1] {
    return true
  }
  
  return false
}

func part1(input [][]string)  { 
  var t []int
  counter := 0

  // Take the input line by line convert each string to an in 
  // then add them to a slice of ints and send them to 
  // checkifvalidTriangle to check if they are valid triangle
  // if they are valid increase the counter by 1
  // final result is the counter
  for _, slice := range input {
    for _, line := range slice {
      s := strings.Fields(line)
      for _, n := range s {
        temp, _ := strconv.Atoi(n)
        t = append(t, temp)
      }
      if checkifvalidTriangle(t) {
        counter += 1
      }
      t = nil
    }
  }
  fmt.Println(counter)
}

// Surely it must be a better and cleaner way to do this
// but this is what I came up on my own at that time.
func part2(input [][]string)  {
  var triangles [1992][3]int
  var t []int
  var transposed [3][1992]int
  counter := 0
  i := 0
  j := 0
  k := 0

  // Similar thing as in part1 but now I convert them to in 
  // I put them in a slice of slices each row is a line and
  // each column is one side of a triangle
  for _, slice := range input {
    for _, line := range slice {
      s := strings.Fields(line)
      for _, n := range s {
        temp, _ := strconv.Atoi(n)
        triangles[i][j] = temp
        j += 1
      }
      j = 0
      i += 1
    }
  }

  // Then I converted the columns to rows and saved it as transposed
  for i=0; i<3; i++ {
    for j=0; j<1992; j++ {
      transposed[i][j] = triangles[j][i]
    }
  }
  // Then I iterated throught the transposed with the a step of 3
  // and I checked those 3 values if they are valid triangles
  // and increaed the counter if they are.
  for i=0; i<3; i++ {
    for j=0; j<1992; j+=3 {
      for k=0; k<3; k++ {
        t = append(t, transposed[i][j+k])
      }
      if checkifvalidTriangle(t) {
        counter += 1
      }
      t = nil
    }
  }
  fmt.Println(counter)
}
