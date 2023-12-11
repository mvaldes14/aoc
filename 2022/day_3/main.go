package main

import (
	"fmt"
	"unicode"
  U "aoc2022"
)


func main() {
	firstPart()
	secondPart()
}

func firstPart() {
	data := U.ReadFile()
	items := make(map[rune]bool)
	sum := 0
	for _, line := range data {
		first, second := line[:len(line)/2], line[len(line)/2:]
		for _, i := range first {
			items[i] = true
		}
		for _, j := range second {
			if items[j] {
				// Clear map on found
				for k := range items {
					delete(items, k)
				}
				// Counts
				if unicode.IsUpper(j) {
					sum += int(j) - 'A' + 27
				} else {
					sum += int(j) - 'a' + 1
				}
			}
		}
	}
  fmt.Println(sum)
}

func secondPart() {
  data := U.ReadFile()
  sum := 0
  for i := 0; i <len(data); i += 3{
    g1,g2,g3 := data[i], data[i+1], data[i+2]
    items1 := createSet(g1)
    items2 := createSet(g2)
    items3 := createSet(g3)
    for line := range items1 {
      if items2[line] && items3[line] {
				if unicode.IsUpper(line) {
					sum += int(line) - 'A' + 27
				} else {
					sum += int(line) - 'a' + 1
				}
      }
    }
  }
  fmt.Println(sum)
}

func createSet(items string)(set map[rune]bool){
  set = make(map[rune]bool)
	for _, item := range items{
		set[item] = true
	}
	return
}

