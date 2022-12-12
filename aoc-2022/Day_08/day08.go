package main

import (
	"bufio"
	"fmt"
	"os"
)

// the coordinate for each tree
type point struct {
  x int;
  y int;
}

// helping method for point summarization
func (p1 *point) Add(p2 point) point {

  return point{p1.x + p2.x, p1.y + p2.y}
}

// helping method for point  for multiplication
func (p1 *point) Mul(n int) point {

  return point{p1.x * n, p1.y * n}
}

func main()  {
  var input []string
  file , _ := os.Open("input.txt")
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
  // create a map for the coords of each tree and their height
  field := map[point]int{}

  for x, line := range input {
    for y, char := range line {
      field[point{x, y}] = int(char - '0')
    }
  }

  vis_count := 0
  dirs := []point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

  // iterate over each tree
  for coord, height := range field {

    // and check for each of the of the 4 directions
    // if the visibility is blocked by another tree
    out:
    for _, d :=  range dirs {
      for i:=1; ; i++ {
        to_cmp, ok := field[coord.Add(d.Mul(i))]
        if !ok {
          vis_count++
          break out // break out of both loops since we only need to have visibility from only one side
        } else if to_cmp >= height {
          break
        }
      }
    }
  }

  fmt.Println(vis_count)
}

func part2(input []string) {
  // Same thing as part 1 but this time calculating score instead of visibility count
  field := map[point]int{}

  for x, line := range input {
    for y, char := range line {
      field[point{x, y}] = int(char - '0')
    }
  }

  best_score := 0
  dirs := []point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

  for coord, height := range field {
    score := 1

    for _, d :=  range dirs {
      for i:=1; ; i++ {
        to_cmp, ok := field[coord.Add(d.Mul(i))]
        if !ok {
          if i == 1 { score = 0} else { score *= i-1 }
          break  // no need to break out of both loops anymore
        } else if to_cmp >= height {
          score *= i
          break
        }
      }
    }
    if score > best_score { best_score = score }
  }

  fmt.Println(best_score)
}
