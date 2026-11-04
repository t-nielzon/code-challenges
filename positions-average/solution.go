package main

import (
	"fmt"
	"strings"
)

func PosAverage(s string) string {
	substrings := strings.Split(s, ", ")
	n := len(substrings)
	
	for i := range substrings {
		substrings[i] = strings.TrimSpace(substrings[i])
	}
	
	totalMatches := 0
	pairCount := 0
	
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			pairCount++
			for k := 0; k < len(substrings[i]); k++ {
				if substrings[i][k] == substrings[j][k] {
					totalMatches++
				}
			}
		}
	}
	
	totalPositions := pairCount * len(substrings[0])
	percentage := float64(totalMatches) / float64(totalPositions) * 100
	
	return fmt.Sprintf("%.10f", percentage)
}