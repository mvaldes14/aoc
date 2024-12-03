// package main for day 2
package main

import (
	U "aoc_2024"
	"fmt"
	"strconv"
	"strings"
)

func checkSafe(num int) bool {
	if num > 3 {
		fmt.Print("Greater than: ", num, "\n")
		return false
	}
	return true
}

func main() {
	input := U.ReadFile()
	safe := 0
	for _, line := range input {
		parts := strings.Fields(line)
		for i := range parts {
			maxLen := len(parts)
			diff := 0
			// NOTE: This is to prevent an index error
			if i+1 >= maxLen {
				continue
			}
			current, _ := strconv.Atoi(parts[i])
			next, _ := strconv.Atoi(parts[i+1])
			fmt.Println(current, next)
			//TODO: Replace with a switch since we have like 4 different cases
			if current > next {
				diff += current - next
				if checkSafe(diff) {
					safe++
				}
				continue
			} else if next > current {
				diff += next - current
				if checkSafe(diff) {
					safe++
				}
				continue
			} else if diff == 0 {
				fmt.Println(current, next)
				fmt.Println("No increase nor decrease")
				continue
			} else if next < diff || next > diff {
				fmt.Println(current, next)
				fmt.Println("Unsafe, increase or decrease")
				continue
			}
			diff = 0
		}
	}
	fmt.Println(safe)

}
