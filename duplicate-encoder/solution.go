package kata

import "strings"

func DuplicateEncoder(word string) string {
	lower := strings.ToLower(word)
	
	count := make(map[rune]int)
	for _, ch := range lower {
		count[ch]++
	}
	
	result := ""
	for _, ch := range lower {
		if count[ch] == 1 {
			result += "("
		} else {
			result += ")"
		}
	}
	
	return result
}