package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func findAnagrams(words []string) map[string][]string {
	anagramMap := make(map[string][]string)
	for _, word := range words {
		sortedWord := sortString(strings.ToLower(word))
		anagramMap[sortedWord] = append(anagramMap[sortedWord], word)
	}
	return anagramMap
}

func main() {
	wordList := []string{"cat", "dog", "tac", "god", "act", "good"}

	anagramGroups := findAnagrams(wordList)

	for _, group := range anagramGroups {
		fmt.Println(group)
	}
}
