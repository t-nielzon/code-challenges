package kata

import (
	"fmt"
	"strings"
)

func EncryptThis(text string) string {
	if text == "" {
		return ""
	}
	words := strings.Split(text, " ")
	result := make([]string, len(words))
	for i, word := range words {
		runes := []rune(word)
		switch len(runes) {
		case 1:
			result[i] = fmt.Sprintf("%d", runes[0])
		case 2:
			result[i] = fmt.Sprintf("%d%c", runes[0], runes[1])
		default:
			runes[1], runes[len(runes)-1] = runes[len(runes)-1], runes[1]
			result[i] = fmt.Sprintf("%d%s", runes[0], string(runes[1:]))
		}
	}
	return strings.Join(result, " ")
}