package kata

import (
	"strings"
	"unicode"
)

func CamelCase(s string) string {
	words := strings.Fields(s)
	var result strings.Builder
	
	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		runes := []rune(word)
		runes[0] = unicode.ToUpper(runes[0])
		result.WriteString(string(runes))
	}
	
	return result.String()
}