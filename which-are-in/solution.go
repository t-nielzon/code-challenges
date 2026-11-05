package kata

import (
	"sort"
	"strings"
)

func InArray(a1 []string, a2 []string) []string {
	result := []string{}
	seen := make(map[string]bool)

	for _, s1 := range a1 {
		for _, s2 := range a2 {
			if strings.Contains(s2, s1) {
				if !seen[s1] {
					result = append(result, s1)
					seen[s1] = true
				}
				break
			}
		}
	}

	sort.Strings(result)
	return result
}