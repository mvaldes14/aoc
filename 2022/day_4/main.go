package main

import (
	U "aoc2022"
	"fmt"
)

func main() {
	data := U.ReadFile()
	contained := 0
	overlap := 0
	for _, line := range data {
		var startFirst, endFirst, startSecond, endSecond int
		fmt.Sscanf(line, "%d-%d,%d-%d", &startFirst, &endFirst, &startSecond, &endSecond)
		fmt.Println(startFirst, endFirst, startSecond, endSecond)
		if startSecond >= startFirst && endSecond <= endFirst || startFirst >= startSecond && endFirst <= endSecond {
			contained++
		}
		if startSecond <= endFirst && endSecond >= startFirst || startFirst <= endSecond && endFirst >= startSecond {
			overlap++
		}
	}
	fmt.Println("first part:", contained)
	fmt.Println("second part:", overlap)
}
