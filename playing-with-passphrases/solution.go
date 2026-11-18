// solution.go
package main

import "strings"

func PassphraseEncoder(text string, shift int) string {
	// Convert to uppercase
	text = strings.ToUpper(text)
	
	// Step 1, 2, 3: shift letters, replace digits, keep non-alphanumeric
	transformed := ""
	for _, ch := range text {
		if ch >= 'A' && ch <= 'Z' {
			// Shift letter with circular wrap
			pos := (((int(ch - 'A') + shift) % 26) + 26) % 26
			shifted := rune(pos) + 'A'
			transformed += string(shifted)
		} else if ch >= '0' && ch <= '9' {
			// Replace digit with complement to 9
			digit := int(ch - '0')
			complement := 9 - digit
			transformed += string(rune(complement) + '0')
		} else {
			// Keep non-alphanumeric and non-digit characters
			transformed += string(ch)
		}
	}
	
	// Step 4: downcase odd positions, upcase even positions
	cased := ""
	for i, ch := range transformed {
		if ch >= 'A' && ch <= 'Z' {
			if i%2 == 0 {
				// Even position: upcase (already uppercase)
				cased += string(ch)
			} else {
				// Odd position: downcase
				cased += string(ch + 32)
			}
		} else {
			// Non-letter characters remain unchanged
			cased += string(ch)
		}
	}
	
	// Step 5: reverse
	runes := []rune(cased)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	
	return string(runes)
}