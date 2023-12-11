package main

import (
	U "aoc2022"
	"fmt"
)

func main(){
  data := U.ReadFile()
  // grid := [][]int{}

  for row := range data{
    for col := range data[row]{
      fmt.Println(data[row], data[col])
    }
  }

}
