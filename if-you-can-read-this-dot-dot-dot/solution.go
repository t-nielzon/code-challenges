package main

import (
	"strings"
	"unicode"
)

func ToNato(words string) string {
	nato := map[rune]string{
		'A': "Alfa", 'B': "Bravo", 'C': "Charlie", 'D': "Delta",
		'E': "Echo", 'F': "Foxtrot", 'G': "Golf", 'H': "Hotel",
		'I': "India", 'J': "Juliett", 'K': "Kilo", 'L': "Lima",
		'M': "Mike", 'N': "November", 'O': "Oscar", 'P': "Papa",
		'Q': "Quebec", 'R': "Romeo", 'S': "Sierra", 'T': "Tango",
		'U': "Uniform", 'V': "Victor", 'W': "Whiskey", 'X': "Xray",
		'Y': "Yankee", 'Z': "Zulu",
	}

	var result []string
	for _, ch := range words {
		if unicode.IsLetter(ch) {
			upperCh := unicode.ToUpper(ch)
			result = append(result, nato[upperCh])
		} else if strings.ContainsAny(string(ch), ",.!?") {
			result = append(result, string(ch))
		}
	}

	return strings.Join(result, " ")
}