package main

import (
	"sort"
	"strings"
)

func InArray(a1, a2 []string) []string {
	result := make([]string, 0)
	seen := make(map[string]bool)
	
	for _, str1 := range a1 {
		if seen[str1] {
			continue
		}
		
		for _, str2 := range a2 {
			if strings.Contains(str2, str1) {
				result = append(result, str1)
				seen[str1] = true
				break
			}
		}
	}
	
	sort.Strings(result)
	return result
}