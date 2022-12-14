package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
  x int
  y int
}

func (p *point) Add(p2 point) point {
  
  return point{p.x + p2.x, p.y + p2.y}
}

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
  // var grid []point
  var visited []point
  head := point{0, 0}
  tail := point{0, 0}
  visited = append(visited, point{0, 0})

  
  for _, line := range input {
    motion := strings.Fields(line)
    // var tmp []point

    // I better way to to do all these must exists, but whatever...
    switch motion[0] {
      case "R": fmt.Println("Right → ", motion[1], "steps")
                steps, _ := strconv.Atoi(motion[1])
                fmt.Println("Steps: ")
                for i:=1; i<=steps; i++ { 
                  head = moveHead(head, point{0 ,1})
                  tail = moveTail(head, tail)
                  visited = append(visited, tail)
                  fmt.Println("H: ", head, "T: ", tail)
                }
      case "U": fmt.Println("Up ↑", motion[1], "steps")
                steps, _ := strconv.Atoi(motion[1])
                fmt.Println("Steps: ")
                for i:=1; i<=steps; i++ {
                  head = moveHead(head, point{-1, 0})
                  tail = moveTail(head, tail)
                  visited = append(visited, tail)
                  fmt.Println("H: ", head, "T: ", tail)
                }
      case "D": fmt.Println("Down ↓", motion[1], "steps")
                steps, _ := strconv.Atoi(motion[1])
                fmt.Println("Steps: ")
                for i:=1; i<=steps; i++ {
                  head = moveHead(head, point{1, 0})    
                  tail = moveTail(head, tail)
                  visited = append(visited, tail)
                  fmt.Println("H: ", head, "T: ", tail)
                }
      case "L": fmt.Println("Left ←", motion[1], "steps")
                steps, _ := strconv.Atoi(motion[1])
                fmt.Println("Steps: ")
                for i:=1; i<=steps; i++ {
                  head = moveHead(head, point{0, -1})  
                  tail = moveTail(head, tail)
                  visited = append(visited, tail)
                  fmt.Println("H: ", head, "T: ", tail)
                }
      default: fmt.Println("Error!") 
    }
    fmt.Println("------------------------------------")
  }
  
  fmt.Println("Tail visited locations: ", visited)

  visit_pos := make(map[point]int)
  for _, v := range visited {
    visit_pos[v] += 1
  }
  fmt.Println("Map of tail pos: ", visit_pos)
  fmt.Println("length of map: ", len(visit_pos))
}

func part2(input []string)  {
  
}

// obviously is moving the head
func moveHead(head point, dir point) point {
  
  return head.Add(dir)
}

// the fuck I just did here ?
// obviously is moving the tail
func moveTail(head point, tail point) point {
  // check if they are adjacent
  if !adjacent(head, tail) {
    // in not check where the head is in relation to the moveTail 
    // and move the tail accordingly
    if head.x == tail.x {
      // if the x are equal then you need to move the y
      dy := head.y - tail.y
      dy = dy / absInt(dy)
      tail.y += dy
    } else if head.y == tail.y {
      // if the y are equal need to move the x
      dx := head.x - tail.x
      dx = dx / absInt(dx)
      tail.x += dx
    } else {
      // if neither are equal mean the are diagonaly away
      dx := head.x - tail.x
      dy := head.y - tail.y
      dx = dx / absInt(dx)
      dy = dy / absInt(dy)
      tail.x += dx
      tail.y += dy
    }
  } 

  return tail
}

// check if they are adjacent
func adjacent(head point, tail point) bool {
  // if their x and y are the same well they overlaping
  if head.x == tail.x && head.y == tail.y {
    return true
  }
  // if their distances are 1 then the are adjacent
  if (absInt(head.x - tail.x) + absInt(head.y - tail.y)) == 1 || absInt(head.x - tail.x) == 1 && absInt(head.y - tail.y) == 1 {
    return true
  }
  
  return false
}

// simple absolute function for integers only
func absInt(n int) int {
  if n > 0 {
    return n
  }

  return -n
}
