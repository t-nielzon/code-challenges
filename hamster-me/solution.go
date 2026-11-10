package main

import "strconv"

func HamsterMe(code, message string) string {
	// Extract unique code letters
	codeSet := make(map[rune]bool)
	for _, ch := range code {
		codeSet[ch] = true
	}

	var codeLetters []rune
	for ch := 'a'; ch <= 'z'; ch++ {
		if codeSet[ch] {
			codeLetters = append(codeLetters, ch)
		}
	}

	// Build cipher mapping for each letter a-z
	cipher := make(map[rune]string)
	for letterIdx := 0; letterIdx < 26; letterIdx++ {
		letter := rune('a' + letterIdx)

		// Find the largest code letter <= letter
		codeBase := codeLetters[len(codeLetters)-1] // default to last code letter
		for _, cb := range codeLetters {
			if cb <= letter {
				codeBase = cb
			}
		}

		// Calculate level as 1-indexed distance from code base letter
		diff := (int(letter) - int(codeBase)) % 26
		if diff < 0 {
			diff += 26
		}
		level := diff + 1

		cipher[letter] = string(codeBase) + strconv.Itoa(level)
	}

	// Encode message using cipher
	var result string
	for _, ch := range message {
		result += cipher[ch]
	}

	return result
}