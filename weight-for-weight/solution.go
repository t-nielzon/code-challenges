package kata

import (
	"sort"
	"strings"
)

func OrderWeight(strng string) string {
	fields := strings.Fields(strng)
	sort.SliceStable(fields, func(i, j int) bool {
		wi, wj := digitSum(fields[i]), digitSum(fields[j])
		if wi != wj {
			return wi < wj
		}
		return fields[i] < fields[j]
	})
	return strings.Join(fields, " ")
}

func digitSum(s string) int {
	sum := 0
	for _, c := range s {
		sum += int(c - '0')
	}
	return sum
}