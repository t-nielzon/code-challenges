package kata

import "strings"

func Arrange(s string) string {
	words := strings.Fields(s)
	// single left-to-right pass: swapping a violating adjacent pair carries
	// the moved word along, exactly matching the described process
	for i := 0; i < len(words)-1; i++ {
		a, b := len(words[i]), len(words[i+1])
		if i%2 == 0 {
			if a > b {
				words[i], words[i+1] = words[i+1], words[i]
			}
		} else {
			if a < b {
				words[i], words[i+1] = words[i+1], words[i]
			}
		}
	}
	for i, w := range words {
		if i%2 == 0 {
			words[i] = strings.ToLower(w)
		} else {
			words[i] = strings.ToUpper(w)
		}
	}
	return strings.Join(words, " ")
}