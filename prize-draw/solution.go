package kata

import (
	"sort"
	"strings"
	"unicode"
)

func Rank(st string, we []int, n int) string {
	if st == "" {
		return "No participants"
	}
	names := strings.Split(st, ",")
	if n > len(names) {
		return "Not enough participants"
	}

	type entry struct {
		name string
		wn   int
	}

	entries := make([]entry, len(names))
	for i, name := range names {
		som := len(name)
		for _, ch := range name {
			som += int(unicode.ToLower(ch)) - int('a') + 1
		}
		entries[i] = entry{name, som * we[i]}
	}

	sort.Slice(entries, func(i, j int) bool {
		if entries[i].wn != entries[j].wn {
			return entries[i].wn > entries[j].wn
		}
		return entries[i].name < entries[j].name
	})

	return entries[n-1].name
}