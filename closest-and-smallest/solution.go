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
	value  int
}

func digitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
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
		val, _ := strconv.Atoi(p)
		nums[i] = numInfo{
			weight: digitSum(val),
			index:  i,
			value:  val,
		}
	}

	// Sort by weight, then by index
	sorted := make([]numInfo, len(nums))
	copy(sorted, nums)
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].weight != sorted[j].weight {
			return sorted[i].weight < sorted[j].weight
		}
		return sorted[i].index < sorted[j].index
	})

	// Find the closest pair
	var bestA, bestB numInfo
	bestDiff := -1

	for i := 0; i < len(sorted)-1; i++ {
		a, b := sorted[i], sorted[i+1]
		diff := b.weight - a.weight

		if bestDiff == -1 || diff < bestDiff {
			bestDiff = diff
			bestA, bestB = a, b
		} else if diff == bestDiff {
			// Compare by smallest weight first
			if a.weight < bestA.weight {
				bestA, bestB = a, b
			} else if a.weight == bestA.weight {
				// Compare by smallest index of first element
				if a.index < bestA.index {
					bestA, bestB = a, b
				} else if a.index == bestA.index {
					// Compare by smallest index of second element
					if b.index < bestB.index {
						bestA, bestB = a, b
					}
				}
			}
		}
	}

	// Format output - ensure sorted by weight, then by index
	first, second := bestA, bestB
	if first.weight > second.weight || (first.weight == second.weight && first.index > second.index) {
		first, second = second, first
	}

	return fmt.Sprintf("[(%d, %d, %d), (%d, %d, %d)]",
		first.weight, first.index, first.value,
		second.weight, second.index, second.value)
}