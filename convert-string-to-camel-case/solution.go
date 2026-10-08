package kata

import "strings"

func ToCamelCase(str string) string {
	if str == "" {
		return str
	}
	
	// Replace underscores with dashes for uniform splitting
	str = strings.ReplaceAll(str, "_", "-")
	
	// Split by dash delimiter
	words := strings.Split(str, "-")
	
	// Capitalize all words after the first one
	for i := 1; i < len(words); i++ {
		if len(words[i]) > 0 {
			words[i] = strings.ToUpper(string(words[i][0])) + strings.ToLower(words[i][1:])
		}
	}
	
	return strings.Join(words, "")
}