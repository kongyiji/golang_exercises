package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	countMap := make(map[string]int)
	wordList := strings.Fields(s)
	l := len(wordList)
	for i := 0; i < l; i++ {
		countMap[wordList[i]]++
	}
	return countMap
}

func main() {
	wc.Test(WordCount)
}
