package kata

import "strings"

func DuplicateEncode(word string) string {
	word = strings.ToLower(word)
	
	freq := make(map[rune]int)
	for _, char := range word {
		freq[char]++
	}
	
	result := make([]rune, 0, len(word))
	for _, char := range word {
		if freq[char] == 1 {
			result = append(result, '(')
		} else {
			result = append(result, ')')
		}
	}
	
	return string(result)
}