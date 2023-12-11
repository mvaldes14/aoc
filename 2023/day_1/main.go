package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	// Open the file
	file, _ := os.Open("inputs.txt")
	results := []rune{}
	total := []string{}

	// Read the contents of the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, char := range line {
			if unicode.IsNumber(char) {
				results = append(results, char)
				continue
			}
		}
		first := results[0]
		last := results[len(results)-1]
    full := string(first) + string(last)
		total = append(total, full )
		results = []rune{}
	}

	final := 0
	for _, char := range total {
    d , _ := strconv.Atoi(char) 
		final += d 
	}

	fmt.Println(final)

	// Close the file
	defer file.Close()

}
