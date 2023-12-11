package main

import (
	U "aoc2022"
	"fmt"
	"strings"
)

func main() {
	// A=Rock, B=Paper, C=Scissors, X=Rock, Y=Paper, Z=Scissors
	// Rock =1, Paper=2, Scissors =3
	// Win=6, Lose=0, Draw=0
	type combinations struct {
		oponent string
		me      string
	}

	combos_pt1 := map[combinations]int{
		{oponent: "A", me: "X"}: 4,
		{oponent: "A", me: "Y"}: 8,
		{oponent: "A", me: "Z"}: 3,
		{oponent: "B", me: "X"}: 1,
		{oponent: "B", me: "Y"}: 5,
		{oponent: "B", me: "Z"}: 9,
		{oponent: "C", me: "X"}: 7,
		{oponent: "C", me: "Y"}: 2,
		{oponent: "C", me: "Z"}: 6,
	}

	combos_pt2 := map[combinations]int{
		{oponent: "A", me: "X"}: 3,
		{oponent: "A", me: "Y"}: 4,
		{oponent: "A", me: "Z"}: 8,
		{oponent: "B", me: "X"}: 1,
		{oponent: "B", me: "Y"}: 5,
		{oponent: "B", me: "Z"}: 9,
		{oponent: "C", me: "X"}: 2,
		{oponent: "C", me: "Y"}: 6,
		{oponent: "C", me: "Z"}: 7,
	}
  score_pt1 := 0
  score_pt2 := 0
  for _, line := range U.ReadFile() {
	l := strings.Fields(line)
	oponent := l[0]
	me := l[1]
    temp := combinations{oponent, me}
    score_pt1 += combos_pt1[temp]
    score_pt2 += combos_pt2[temp]

  }
  fmt.Println(score_pt1)
  fmt.Println(score_pt2)
}
