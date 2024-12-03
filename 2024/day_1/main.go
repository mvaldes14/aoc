package main

import (
	U "aoc_2024"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func makeLists() ([]int, []int) {
	input := U.ReadFile()
	var left, right []int
	left = make([]int, len(input))
	right = make([]int, len(input))
	for i, val := range input {
		split := strings.Fields(val)
		left[i], _ = strconv.Atoi(split[0])
		right[i], _ = strconv.Atoi(split[1])
	}
	return left, right
}

func partOne() {
	left, right := makeLists()
	sort.Ints(left)
	sort.Ints(right)
	total := 0
	for i := range left {
		dif := 0
		if left[i] < right[i] {
			dif = left[i] - right[i]
		} else {
			dif = right[i] - left[i]
		}
		total += dif
	}
	fmt.Println(total * -1)
}

func partTwo() {
	left, right := makeLists()
	similar := 0
	listDups := make(map[int]int)
	sort.Ints(left)
	sort.Ints(right)
	for _, v := range left {
		listDups[v]++
	}
	for _, v := range right {
		similar += v * listDups[v]
	}
	fmt.Println(similar)

}
func main() {
	partOne()
	partTwo()
}
