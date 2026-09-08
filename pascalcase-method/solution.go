package main

import "strings"

func Camelcase(str string) string {
	words := strings.Fields(str)
	for i := range words {
		if len(words[i]) > 0 {
			words[i] = strings.ToUpper(string(words[i][0])) + strings.ToLower(words[i][1:])
		}
	}
	return strings.Join(words, "")
}