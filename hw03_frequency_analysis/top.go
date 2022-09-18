package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(str string) []string {
	wordList := strings.Fields(str)
	counts := make(map[string]int)
	for _, word := range wordList {
		counts[word]++
	}

	words := make([]string, 0, len(wordList))
	for word := range counts {
		words = append(words, word)
	}

	sort.Slice(words, func(i, j int) bool {
		if counts[words[i]] == counts[words[j]] {
			return words[i] < words[j]
		}
		return counts[words[i]] > counts[words[j]]
	})

	resultLimit := 10
	if len(words) < resultLimit {
		resultLimit = len(words)
	}
	return words[:resultLimit]
}
