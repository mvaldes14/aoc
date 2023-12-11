package main

import (
	U "aoc2022"
	"fmt"
	"strings"
)

func main() {
	data := U.ReadFile()
	var currentFolder string
	topfolders := make(map[string]int)
	var folderStack []string
	for _, line := range data {
		var filename string
		var filesize int
		fmt.Sscanf(line, "%d %s", &filesize, &filename)
		topfolders[currentFolder] += filesize
		if strings.Contains(line, "cd") {
			currentFolder = strings.Split(line, " ")[2]
			folderStack = append(folderStack, currentFolder)
			if currentFolder == ".." {
        folderStack = folderStack[:len(folderStack)-2]
        continue
			}
		}
	}
  fmt.Println(folderStack)
	fmt.Println(topfolders)
}
