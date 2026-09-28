package main

import (
	"sort"
	"strconv"
	"strings"
)

func HamsterMe(code string, message string) string {
	// Extract unique code letters and sort them
	codeSet := make(map[rune]bool)
	for _, ch := range code {
		codeSet[ch] = true
	}
	
	var codeLetters []rune
	for ch := range codeSet {
		codeLetters = append(codeLetters, ch)
	}
	sort.Slice(codeLetters, func(i, j int) bool {
		return codeLetters[i] < codeLetters[j]
	})
	
	var result strings.Builder
	
	for _, char := range message {
		// Find the closest preceding (or equal) code letter
		var closest rune
		found := false
		
		for i := len(codeLetters) - 1; i >= 0; i-- {
			if codeLetters[i] <= char {
				closest = codeLetters[i]
				found = true
				break
			}
		}
		
		// If no code letter found, wrap to the last one
		if !found {
			closest = codeLetters[len(codeLetters)-1]
		}
		
		// Calculate distance and row number
		var distance int
		if char >= closest {
			distance = int(char) - int(closest)
		} else {
			distance = 26 - int(closest) + int(char)
		}
		
		row := distance + 1
		result.WriteRune(closest)
		result.WriteString(strconv.Itoa(row))
	}
	
	return result.String()
}