package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"sort"
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
  files := map[string]int{} // to hold all the directories and their total size
  cd := "" // path name is taken by the path module :P
  sum := 0

  for _, line := range input {
    var size int
    var name string

    // For each line check if it's a command, if it is
    // use path module to create the file path
    // else if the format %number %string exist
    // for every directory in cd remove the last directory and
    // add the size of it's files in the files map[directory]size 
    // untill we are at root "/"
    if strings.HasPrefix(line, "$ cd") {
      cd = path.Join(cd, strings.Fields(line)[2])
    } else if _, err := fmt.Sscanf(line, "%d %s", &size, &name); err == nil {
      for d:=cd; d!="/"; d=path.Dir(d) { // path.Dir(d) return everything but the last directory of d
        files[d] += size
      }
      files["/"] += size // sum them all at root
    }
  }

  // sum the files that have 100000 or less size
  for _, f := range files {
    if f <= 100000 {
      sum += f
    }
  }

  fmt.Println(sum)
}

func part2(input []string)  {
  files := map[string]int{} // to hold all the directories and their total size
  cd := "" // path name is taken by the path module :P
  total_storage := 70000000  // total device storage
  needed_storage := 30000000 // how much storage needed to update

  for _, line := range input {
    var size int
    var name string

    if strings.HasPrefix(line, "$ cd") {
      cd = path.Join(cd, strings.Fields(line)[2])
    } else if _, err := fmt.Sscanf(line, "%d %s", &size, &name); err == nil {
      for d:=cd; d!="/"; d=path.Dir(d) { // path.Dir(d) return everything but the last directory of d
        files[d] += size
      }
      files["/"] += size // sum them all at root
    }
  }
  // until here is the same as part 1
  
  // find out the free storage and how much it's needed to 
  // free to be able to updated
  free_storage := total_storage - files["/"]
  storage_needed_for_update := needed_storage - free_storage

  // make a list of the values of all those directories
  // that have enough space to be worth deleting
  var values []int
  for k, v := range files {
    if files[k] >= storage_needed_for_update {
      values = append(values, v)
    }
  }
  sort.Ints(values) // sort them
  fmt.Println(values[0]) // and take the min
  // since we don't need the key this will be enough
}

