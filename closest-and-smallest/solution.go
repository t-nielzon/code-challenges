package kata

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// triple holds the weight, original index, and value of a number.
type triple struct {
	weight int
	index  int
	value  int
}

func digitWeight(n int) int {
	w := 0
	for n > 0 {
		w += n % 10
		n /= 10
	}
	return w
}

func closest(strng string) string {
	fields := strings.Fields(strng)
	if len(fields) == 0 {
		return "[(), ()]"
	}

	triples := make([]triple, 0, len(fields))
	for i, f := range fields {
		v, _ := strconv.Atoi(f)
		triples = append(triples, triple{digitWeight(v), i, v})
	}

	// sort by weight, then by original index, so adjacent pairs are the
	// closest candidates and ties resolve to the smallest weights/indices.
	sort.Slice(triples, func(i, j int) bool {
		if triples[i].weight != triples[j].weight {
			return triples[i].weight < triples[j].weight
		}
		return triples[i].index < triples[j].index
	})

	bestDiff := -1
	var a, b triple
	for i := 0; i+1 < len(triples); i++ {
		diff := triples[i+1].weight - triples[i].weight
		if bestDiff == -1 || diff < bestDiff {
			bestDiff = diff
			a, b = triples[i], triples[i+1]
		}
	}

	return fmt.Sprintf("[(%d, %d, %d), (%d, %d, %d)]",
		a.weight, a.index, a.value, b.weight, b.index, b.value)
}