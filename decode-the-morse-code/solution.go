package kata

import "strings"

func DecodeMorse(morseCode string) string {
	morseCode = strings.TrimSpace(morseCode)
	words := strings.Split(morseCode, "   ")
	decoded := make([]string, len(words))
	for i, word := range words {
		letters := strings.Split(strings.TrimSpace(word), " ")
		for _, letter := range letters {
			if letter != "" {
				decoded[i] += MORSE_CODE[letter]
			}
		}
	}
	return strings.Join(decoded, " ")
}