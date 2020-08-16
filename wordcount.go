package main

import (
	"fmt"
	"strings"
)

func wordcount(str string) map[string]int {

	wordList := strings.Fields(str)
	counts := make(map[string]int)

	for _, word := range wordList {
		_, yes := counts[word]
		if yes {
			counts[word]++
		} else {
			counts[word] = 1
		}
	}
	return counts
}

func main() {
	strLine := "text type here your text 404"

	for index, element := range wordcount(strLine) {
		fmt.Println(index, element)
	}
}
