package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func calculate() []int {
  data,err := os.Open("./input.txt")
  if err != nil {
    panic(err) }
  defer data.Close()

	var ( 
    txtlines []string
    calories []int
  )

  scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
 
  counter := 0
	for _, line := range txtlines {
    num, err := strconv.Atoi(line)
    counter += num
    if err != nil {
      calories = append(calories, counter)
      counter = 0
    }
	}
  sort.Ints(calories)
  return calories
}


func first_part(calories []int) int{
  return calories[len(calories)-1]
}

func second_part(calories []int) int {
  top3 := calories[len(calories)-3:]
  counter := 0
  for _,v := range top3 {
    counter += v
  }
  return counter
}

func main(){
  fmt.Printf("First Part results are: %d \n", first_part(calculate()))
  fmt.Printf("Second Part results are: %d \n", second_part(calculate()))
}


