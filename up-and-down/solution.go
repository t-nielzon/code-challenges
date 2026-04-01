package kata

import "strings"

func Arrange(s string) string {
	if s == "" {
		return ""
	}
	words := strings.Fields(s)
	n := len(words)

	for i := 0; i < n-1; i++ {
		if i%2 == 0 {
			// even index: length should be <= next
			if len(words[i]) > len(words[i+1]) {
				words[i], words[i+1] = words[i+1], words[i]
			}
		} else {
			// odd index: length should be >= next
			if len(words[i]) < len(words[i+1]) {
				words[i], words[i+1] = words[i+1], words[i]
			}
		}
	}

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			words[i] = strings.ToLower(words[i])
		} else {
			words[i] = strings.ToUpper(words[i])
		}
	}

	return strings.Join(words, " ")
}