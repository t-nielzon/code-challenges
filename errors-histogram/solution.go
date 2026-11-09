package main

import (
	"fmt"
	"sort"
	"strings"
)

func Hist(s string) string {
	// Count error letters
	errors := make(map[rune]int)
	for _, ch := range s {
		if ch == 'u' || ch == 'w' || ch == 'x' || ch == 'z' {
			errors[ch]++
		}
	}
	
	// Sort error letters alphabetically
	var letters []rune
	for ch := range errors {
		letters = append(letters, ch)
	}
	sort.Slice(letters, func(i, j int) bool {
		return letters[i] < letters[j]
	})
	
	// Build histogram
	var result []string
	for _, ch := range letters {
		count := errors[ch]
		countStr := fmt.Sprintf("%d", count)
		// Pad count to 6-character field
		padding := 6 - len(countStr)
		if padding < 0 {
			padding = 0
		}
		line := fmt.Sprintf("%s  %s%s%s", string(ch), countStr, strings.Repeat(" ", padding), strings.Repeat("*", count))
		result = append(result, line)
	}
	
	return strings.Join(result, "\r")
}