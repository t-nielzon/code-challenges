package main

import (
	"sort"
	"strings"
)

func InArray(a1 []string, a2 []string) []string {
	seen := make(map[string]bool)

	for _, s1 := range a1 {
		for _, s2 := range a2 {
			if strings.Contains(s2, s1) {
				seen[s1] = true
				break
			}
		}
	}

	var result []string
	for k := range seen {
		result = append(result, k)
	}

	sort.Strings(result)
	return result
}