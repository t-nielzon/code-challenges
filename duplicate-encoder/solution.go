package main

import "strings"

func DuplicateEncoder(word string) string {
	word = strings.ToLower(word)
	
	freq := make(map[rune]int)
	for _, char := range word {
		freq[char]++
	}
	
	var result strings.Builder
	for _, char := range word {
		if freq[char] == 1 {
			result.WriteRune('(')
		} else {
			result.WriteRune(')')
		}
	}
	
	return result.String()
}