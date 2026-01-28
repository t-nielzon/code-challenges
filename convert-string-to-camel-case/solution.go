package kata

import "strings"

func ToCamelCase(s string) string {
	if s == "" {
		return ""
	}
	s = strings.ReplaceAll(s, "_", "-")
	words := strings.Split(s, "-")
	for i := 1; i < len(words); i++ {
		if len(words[i]) > 0 {
			words[i] = strings.ToUpper(words[i][:1]) + words[i][1:]
		}
	}
	return strings.Join(words, "")
}