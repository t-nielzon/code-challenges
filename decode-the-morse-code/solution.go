package kata

import "strings"

func DecodeMorse(morseCode string) string {
	// Trim leading and trailing spaces
	morseCode = strings.TrimSpace(morseCode)
	
	// Split by triple space to get words
	words := strings.Split(morseCode, "   ")
	
	var result []string
	
	for _, word := range words {
		// Split each word by single space to get individual morse characters
		chars := strings.Split(word, " ")
		var decodedWord strings.Builder
		
		for _, char := range chars {
			if char != "" {
				// MORSE_CODE is a preloaded map provided by Codewars
				if decoded, ok := MORSE_CODE[char]; ok {
					decodedWord.WriteString(decoded)
				}
			}
		}
		
		if decodedWord.Len() > 0 {
			result = append(result, decodedWord.String())
		}
	}
	
	return strings.Join(result, " ")
}