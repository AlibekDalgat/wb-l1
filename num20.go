package main

import (
	"fmt"
	"strings"
)

func reverseWords(input string) string {
	words := strings.Fields(input)
	reversedWords := make([]string, len(words))
	size := len(words)
	for i := range words {
		reversedWords[i] = words[size-i-1]
	}
	return strings.Join(reversedWords, " ")
}

func main() {
	input := "snow dog sun"
	reversed := reverseWords(input)
	fmt.Println(reversed)
}
