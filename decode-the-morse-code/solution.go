package kata

import "strings"

func DecodeMorse(morseCode string) string {
	morseCode = strings.TrimSpace(morseCode)
	words := strings.Split(morseCode, "   ")
	decodedWords := make([]string, 0, len(words))
	for _, word := range words {
		letters := strings.Split(word, " ")
		var sb strings.Builder
		for _, letter := range letters {
			if letter == "" {
				continue
			}
			sb.WriteString(MORSE_CODE[letter])
		}
		decodedWords = append(decodedWords, sb.String())
	}
	return strings.Join(decodedWords, " ")
}