package kata

import (
	"sort"
	"strings"
)

func InArray(array1 []string, array2 []string) []string {
	var r []string
	seen := make(map[string]bool)
	for _, s1 := range array1 {
		if seen[s1] {
			continue
		}
		for _, s2 := range array2 {
			if strings.Contains(s2, s1) {
				r = append(r, s1)
				seen[s1] = true
				break
			}
		}
	}
	sort.Strings(r)
	return r
}