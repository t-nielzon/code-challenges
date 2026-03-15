package kata

import (
	"fmt"
	"sort"
	"strings"
)

func Closest(strng string) string {
	if strng == "" {
		return "[(), ()]"
	}
	parts := strings.Fields(strng)
	n := len(parts)
	if n < 2 {
		return "[(), ()]"
	}

	type entry struct {
		weight int
		index  int
		num    string
	}

	entries := make([]entry, n)
	for i, p := range parts {
		w := 0
		for _, c := range p {
			w += int(c - '0')
		}
		entries[i] = entry{w, i, p}
	}

	// sort by weight, then by index
	sorted := make([]entry, n)
	copy(sorted, entries)
	sort.SliceStable(sorted, func(i, j int) bool {
		if sorted[i].weight != sorted[j].weight {
			return sorted[i].weight < sorted[j].weight
		}
		return sorted[i].index < sorted[j].index
	})

	// find the closest pair among consecutive elements in sorted order
	bestDiff := int(^uint(0) >> 1)
	var bestA, bestB entry

	for i := 0; i < len(sorted)-1; i++ {
		a, b := sorted[i], sorted[i+1]
		diff := b.weight - a.weight
		if diff < bestDiff ||
			(diff == bestDiff && a.weight < bestA.weight) ||
			(diff == bestDiff && a.weight == bestA.weight && a.index < bestA.index) ||
			(diff == bestDiff && a.weight == bestA.weight && a.index == bestA.index && b.index < bestB.index) {
			bestDiff = diff
			bestA = a
			bestB = b
		}
	}

	return fmt.Sprintf("[(%d, %d, %s), (%d, %d, %s)]",
		bestA.weight, bestA.index, bestA.num,
		bestB.weight, bestB.index, bestB.num)
}