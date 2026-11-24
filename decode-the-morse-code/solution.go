package main

import (
	"strings"
)

func DecodeMorse(morse string) string {
	result := ""
	words := strings.Split(strings.TrimSpace(morse), "   ")
	
	for i, word := range words {
		if i > 0 {
			result += " "
		}
		codes := strings.Split(word, " ")
		for _, code := range codes {
			if code != "" {
				if char, ok := MORSE_CODE[code]; ok {
					result += char
				}
			}
		}
	}
	
	return result
}