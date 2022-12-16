package main

import (
	"bufio"
	"fmt"
	"math"
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

  fmt.Println("-=Part 1=-")
  part1(input)
  fmt.Println("-=Part 2=-")
  part2(input)
}

func part1(input []string)  {
  Xmap := make(map[int]int)
  cycle := 1
  X := 1
  n := 0
  sum := 0

  // pretty dumb way to do it but its simple
  // Take each command
  // if its noop just increase the cycle and add the cycle and the value of register X in a map
  // if its addx convert the second string to an int and increase cycle add the cylce:value to the map
  // but do not do the operation yet increase the cycle again then do the operation and add it to the map
  for _, line := range input {
    cmd := strings.Fields(line)
    if cmd[0] == "noop" {
      // do nothing
      cycle++
      Xmap[cycle] = X
    }
    if cmd[0] == "addx" {
      n, _ = strconv.Atoi(cmd[1])
      cycle++
      Xmap[cycle] = X
      cycle++
      X += n
      Xmap[cycle] = X
    }
  }
  // after finishing with all that
  // we iterate from 20 with step 40 and do the calculation for the signal then add it to a sum
  for i:=20; i<=220; i+=40 {
    sum += Xmap[i] * i
  }

  fmt.Println("Sum of singals:", sum)
}

func part2(input []string) {
  // this is a lot harder
  // after some time trying by myself 
  // I decide to look around for other peoples solution
  // I adapted someone else idea here
  var screen [6][40]string
  cycle := 0
  X := 1
  n := 0

  // this part ain't difficult but you use the execute function once for noop
  // and twiche for addx the execute() function take the X register, the current cycle and
  // the screen variable as input (all of them as addreses so you don't make a mess with returning vars)
  for _, line := range input {
    op := strings.Fields(line)

    if op[0] == "noop" {
      execute(&X, &cycle, &n, &screen)
    } else if op[0] == "addx" { 
      execute(&X, &cycle, &n, &screen)
      execute(&X, &cycle, &n, &screen)

      n, _ = strconv.Atoi(op[1])
    }
  }

  // just prints the screen pretty much a 2D slice with 6 row and 40 columns as the screen
  for _, row := range screen {
    for _, col := range row {
      fmt.Print(col)
    }
    fmt.Print("\n")
  }
}

// Create a new func that executes the operations on tick each time you run it
func execute(X *int, cycle *int, n *int, screen *[6][40]string) {
  // Do the addition first and then 
  // zero the n argument
  *X += *n
  *n = 0

  // ok I think I understand it now
  // so *cyle modulo 40 gives us the column we are and
  // the X - col gives us the diference and compare it if it's equal or less then 1 is
  // the same as checking if your current col is between x-1 and x+1
  // then you take the *cycle / 40 that gives you the current row and the *cycle % 40 that gives
  // you the current col and you put # if you are new the X+/-1 you turn on the pixel #
  // and if you are not you turn it off . 
  // thanks https://github.com/17ms :) 
  if math.Abs(float64(*X - (*cycle % 40))) <= 1 {
    (*screen)[*cycle / 40][*cycle % 40] = "#"
  } else {
    (*screen)[*cycle / 40][*cycle % 40] = "."
  }

  // don't forget to increase the cycle count!
  *cycle++
}
