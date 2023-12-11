package utils

import (
	"bufio"
	"os"
)

func ReadFile() []string {
	data, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer data.Close()

	var (
		txtlines []string
	)

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	return txtlines
}

