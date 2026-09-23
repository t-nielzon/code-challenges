package main

import (
	"strings"
	"unicode"
)

func ToCamelCase(str string) string {
	if str == "" {
		return str
	}

	// Replace underscores with dashes to have a single delimiter
	str = strings.ReplaceAll(str, "_", "-")

	// Split by dashes
	words := strings.Split(str, "-")

	// Capitalize the first letter of all words except the first one
	for i := 1; i < len(words); i++ {
		if len(words[i]) > 0 {
			words[i] = string(unicode.ToUpper(rune(words[i][0]))) + words[i][1:]
		}
	}

	return strings.Join(words, "")
}