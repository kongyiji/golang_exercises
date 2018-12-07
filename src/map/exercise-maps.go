package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	// length := len(s)
	countMap = make(map[string]int)
	wordList := strings.Fields(s)
	l := len(wordList)
	for i := 0; i < l; i++ {

	}
	fmt.Printf("%d, %q\n", l, wordList)
	return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
