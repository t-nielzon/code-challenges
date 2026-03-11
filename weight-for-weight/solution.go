package kata

import (
	"sort"
	"strings"
)

func digitWeight(s string) int {
	w := 0
	for _, c := range s {
		w += int(c - '0')
	}
	return w
}

func OrderWeight(strng string) string {
	fields := strings.Fields(strng)
	sort.SliceStable(fields, func(i, j int) bool {
		wi, wj := digitWeight(fields[i]), digitWeight(fields[j])
		if wi != wj {
			return wi < wj
		}
		return fields[i] < fields[j]
	})
	return strings.Join(fields, " ")
}