package kata

import (
	"sort"
	"strings"
)

func OrderWeight(strng string) string {
	parts := strings.Fields(strng)
	sort.SliceStable(parts, func(i, j int) bool {
		wi, wj := digitSum(parts[i]), digitSum(parts[j])
		if wi != wj {
			return wi < wj
		}
		return parts[i] < parts[j]
	})
	return strings.Join(parts, " ")
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