package kata

import (
	"strings"
	"unicode"
)

func toWeirdCase(str string) string {
	words := strings.Split(str, " ")
	for i, word := range words {
		result := make([]rune, len([]rune(word)))
		for j, ch := range word {
			if j%2 == 0 {
				result[j] = unicode.ToUpper(ch)
			} else {
				result[j] = unicode.ToLower(ch)
			}
		}
		words[i] = string(result)
	}
	return strings.Join(words, " ")
}