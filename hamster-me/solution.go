// solution.go
package main

import (
	"sort"
	"strconv"
	"strings"
)

func hamsterMe(code string, message string) string {
	// Extract unique letters from code and sort them
	codeMap := make(map[rune]bool)
	for _, ch := range code {
		codeMap[ch] = true
	}

	var codeLetters []rune
	for ch := range codeMap {
		codeLetters = append(codeLetters, ch)
	}
	sort.Slice(codeLetters, func(i, j int) bool {
		return codeLetters[i] < codeLetters[j]
	})

	// Build mapping using BFS expansion
	type Mapping struct {
		codeLetter rune
		distance   int
	}
	mapping := make(map[rune]Mapping)

	// Initialize code letters at distance 1
	currentLevel := make(map[rune]rune)
	for _, ch := range codeLetters {
		currentLevel[ch] = ch
		mapping[ch] = Mapping{ch, 1}
	}

	// Expand level by level until all 26 letters are assigned
	distance := 1
	for len(mapping) < 26 {
		nextLevel := make(map[rune]rune)
		for letter, codeLetterOrigin := range currentLevel {
			nextLetter := rune(((int(letter) - int('a') + 1) % 26) + int('a'))
			if _, exists := mapping[nextLetter]; !exists {
				nextLevel[nextLetter] = codeLetterOrigin
				mapping[nextLetter] = Mapping{codeLetterOrigin, distance + 1}
			}
		}
		currentLevel = nextLevel
		distance++
	}

	// Encode the message
	var result strings.Builder
	for _, ch := range message {
		m := mapping[ch]
		result.WriteRune(m.codeLetter)
		result.WriteString(strconv.Itoa(m.distance))
	}

	return result.String()
}