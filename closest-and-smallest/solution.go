package kata

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type numInfo struct {
	weight int
	index  int
	value  string
}

func Closest(strng string) string {
	if strng == "" {
		return "[(), ()]"
	}

	parts := strings.Fields(strng)
	if len(parts) < 2 {
		return "[(), ()]"
	}

	nums := make([]numInfo, len(parts))
	for i, p := range parts {
		nums[i] = numInfo{
			weight: digitSum(p),
			index:  i,
			value:  p,
		}
	}

	// sort by weight, then by index
	sorted := make([]numInfo, len(nums))
	copy(sorted, nums)
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].weight != sorted[j].weight {
			return sorted[i].weight < sorted[j].weight
		}
		return sorted[i].index < sorted[j].index
	})

	// find the closest pair by comparing adjacent elements in sorted order
	var bestA, bestB numInfo
	bestDiff := -1

	for i := 0; i < len(sorted)-1; i++ {
		a, b := sorted[i], sorted[i+1]
		diff := b.weight - a.weight

		if bestDiff == -1 || diff < bestDiff {
			bestDiff = diff
			bestA, bestB = a, b
		} else if diff == bestDiff {
			// compare by smallest weight first
			if a.weight < bestA.weight {
				bestA, bestB = a, b
			} else if a.weight == bestA.weight {
				// compare by smallest index of first element
				if a.index < bestA.index {
					bestA, bestB = a, b
				} else if a.index == bestA.index {
					// compare by smallest index of second element
					if b.index < bestB.index {
						bestA, bestB = a, b
					}
				}
			}
		}
	}

	// ensure bestA comes before bestB (by weight, then by index)
	if bestA.weight > bestB.weight || (bestA.weight == bestB.weight && bestA.index > bestB.index) {
		bestA, bestB = bestB, bestA
	}

	valA, _ := strconv.Atoi(bestA.value)
	valB, _ := strconv.Atoi(bestB.value)

	return fmt.Sprintf("[(%d, %d, %d), (%d, %d, %d)]",
		bestA.weight, bestA.index, valA,
		bestB.weight, bestB.index, valB)
}

func digitSum(s string) int {
	sum := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			sum += int(c - '0')
		}
	}
	return sum
}