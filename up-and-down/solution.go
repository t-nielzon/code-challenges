package kata

import "strings"

func Arrange(s string) string {
	if s == "" {
		return s
	}
	
	words := strings.Fields(s)
	n := len(words)
	
	if n == 1 {
		return strings.ToLower(words[0])
	}
	
	// process each position from left to right
	for i := 1; i < n; i++ {
		if i%2 == 1 {  // odd index: need length[i] >= length[i-1]
			if len(words[i]) < len(words[i-1]) {
				words[i], words[i-1] = words[i-1], words[i]
			}
		} else {  // even index: need length[i] <= length[i-1]
			if len(words[i]) > len(words[i-1]) {
				words[i], words[i-1] = words[i-1], words[i]
			}
		}
	}
	
	// format: even indices lowercase, odd indices uppercase
	result := make([]string, n)
	for i, word := range words {
		if i%2 == 0 {
			result[i] = strings.ToLower(word)
		} else {
			result[i] = strings.ToUpper(word)
		}
	}
	
	return strings.Join(result, " ")
}