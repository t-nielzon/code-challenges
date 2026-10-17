package main

import (
	"strings"
	"unicode/utf8"
)

func arrange(s string) string {
	if s == "" {
		return ""
	}

	words := strings.Fields(s)
	n := len(words)

	if n == 1 {
		return strings.ToLower(words[0])
	}

	// bubble sort with up-down constraint
	for {
		satisfied := true
		for i := 0; i < n-1; i++ {
			lenI := utf8.RuneCountInString(words[i])
			lenI1 := utf8.RuneCountInString(words[i+1])

			needLessOrEqual := i%2 == 0

			if (needLessOrEqual && lenI > lenI1) || (!needLessOrEqual && lenI < lenI1) {
				words[i], words[i+1] = words[i+1], words[i]
				satisfied = false
				break
			}
		}

		if satisfied {
			break
		}
	}

	// apply formatting
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			words[i] = strings.ToLower(words[i])
		} else {
			words[i] = strings.ToUpper(words[i])
		}
	}

	return strings.Join(words, " ")
}