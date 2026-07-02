package kata

import "strings"

func PascalCase(str string) string {
	words := strings.Fields(str)
	for i, w := range words {
		words[i] = strings.ToUpper(w[:1]) + w[1:]
	}
	return strings.Join(words, "")
}