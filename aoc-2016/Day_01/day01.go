package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type coords struct {
  x int
  y int
}

func main()  {
  // Get input as a single string and then split it at commas and trim whitespace
  input, _ := ioutil.ReadFile("input.txt")
  instructions := string(input)
  steps := strings.Split(strings.TrimSpace(instructions), ", ")
  // steps := []string{"R8", "R4", "R4", "R8"}
  fmt.Println("steps: ", steps)

  var xy = coords{0, 0}
  pos := []coords{xy}
  heading := 0
  var ans2 []coords
  for _, e := range steps {
    switch e[0] {
      case 'L': heading = mod((heading-1), 4)
      case 'R': heading = mod((heading+1), 4)
      default: fmt.Println("Heading ERROR!"); os.Exit(1)
    }
    fmt.Println("heading: ", heading)
  
    movement, _ := strconv.Atoi(e[1:])
    fmt.Println("move: ", movement)
    for i:=1; i<=movement; i++ {
      switch heading {
        case 0: xy.y += 1
        case 1: xy.x += 1
        case 2: xy.y -= 1
        case 3: xy.x -= 1
        default: fmt.Println("Movement ERROR!"); os.Exit(2)
      }
    }
    if Exists(pos, xy) {
      // fmt.Println("\t\t-=PART 2 ANS: ", abs(xy.x)+abs(xy.y), "=-")
      ans2 = append(ans2, xy)
    }
    fmt.Println("Current possition: ", xy)
    pos = append(pos, xy)
  }

  // freq := make(map[coords]int)
  // for _, e := range pos {
  //   freq[e] += 1
  // }
  //
  // var ans coords
  // for k, v := range freq {
  //   if v == 2 {
  //     ans = k
  //   }
  // }

  fmt.Println("Poss: \n", pos)
  // fmt.Println("Freq: \n", freq)
  fmt.Println("\t-=Day-01_Part-1=-")
  fmt.Println("ANS: ", abs(pos[len(pos)-1].x) + abs(pos[len(pos)-1].y))

  fmt.Println("\t-=Day-01_Part-2=-")
  // fmt.Println("ANS: ", abs(ans.x) + abs(ans.y))
  fmt.Println("ANS: ", abs(ans2[0].x) + abs(ans2[0].y))
  fmt.Println(ans2)
}


func Exists(pos []coords, xy coords) bool {
  for _,e := range pos {
    if (e.x == xy.x) && (e.y == xy.y) {
      return true
    }
  }

  return false
}

// Since there is not abs() function for integers in GO I'll do myself!
func abs(x int) int {
  if x < 0 {
    return -x
  }

  return x
}

// since there is not modulus operand or function in GO I'll do it myself...
func mod(a int, b int) int {
  var res int = a % b
  if ((res < 0 && b > 0) || (res > 0 && b < 0)) {
    return res + b
  }

  return res
}
