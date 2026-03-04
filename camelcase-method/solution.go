package kata

import "strings"

func CamelCase(s string) string {
	var b strings.Builder
	for _, word := range strings.Fields(s) {
		b.WriteString(strings.ToUpper(word[:1]) + word[1:])
	}
	return b.String()
}