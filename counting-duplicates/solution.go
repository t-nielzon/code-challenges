package kata

import "strings"

func duplicateCount(s string) int {
	s = strings.ToLower(s)
	counts := make(map[rune]int)
	for _, c := range s {
		counts[c]++
	}
	result := 0
	for _, n := range counts {
		if n > 1 {
			result++
		}
	}
	return result
}