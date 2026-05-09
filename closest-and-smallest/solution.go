package kata

import (
	"fmt"
	"sort"
	"strings"
)

func Closest(strng string) string {
	if strings.TrimSpace(strng) == "" {
		return "[(), ()]"
	}

	parts := strings.Fields(strng)
	type item struct {
		weight int
		index  int
		number string
	}
	items := make([]item, len(parts))
	for i, p := range parts {
		w := 0
		for _, c := range p {
			if c >= '0' && c <= '9' {
				w += int(c - '0')
			}
		}
		items[i] = item{w, i, p}
	}

	sort.SliceStable(items, func(i, j int) bool {
		if items[i].weight != items[j].weight {
			return items[i].weight < items[j].weight
		}
		return items[i].index < items[j].index
	})

	bestDiff := -1
	bestA, bestB := 0, 1
	for i := 0; i < len(items)-1; i++ {
		d := items[i+1].weight - items[i].weight
		if bestDiff == -1 || d < bestDiff {
			bestDiff = d
			bestA, bestB = i, i+1
		}
	}

	a, b := items[bestA], items[bestB]
	return fmt.Sprintf("[(%d, %d, %s), (%d, %d, %s)]",
		a.weight, a.index, a.number,
		b.weight, b.index, b.number)
}