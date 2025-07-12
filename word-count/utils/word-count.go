package utils

import (
	"strings"
)

func WordCount(str string) map[string]int {
	words := strings.Fields(str)

	wordsWithCounts := map[string]int{}

	for _, word := range words {
		wordsWithCounts[word] = len(word)
	}

	return wordsWithCounts
}