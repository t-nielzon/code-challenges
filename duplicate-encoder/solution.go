package kata

import "strings"

func DuplicateEncode(word string) string {
	lower := strings.ToLower(word)
	counts := make(map[rune]int)
	for _, r := range lower {
		counts[r]++
	}
	var b strings.Builder
	b.Grow(len(lower))
	for _, r := range lower {
		if counts[r] > 1 {
			b.WriteByte(')')
		} else {
			b.WriteByte('(')
		}
	}
	return b.String()
}