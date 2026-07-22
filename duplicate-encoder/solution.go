package main

import "strings"

func DuplicateEncode(word string) string {
	word = strings.ToLower(word)
	
	freq := make(map[rune]int)
	for _, ch := range word {
		freq[ch]++
	}
	
	result := ""
	for _, ch := range word {
		if freq[ch] == 1 {
			result += "("
		} else {
			result += ")"
		}
	}
	
	return result
}