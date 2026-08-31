package main

import (
	"fmt"
	"strings"
)

func PosAverage(s string) string {
	parts := strings.Split(s, ",")
	var substrings []string
	for _, part := range parts {
		substrings = append(substrings, strings.TrimSpace(part))
	}

	n := len(substrings)
	totalCommon := 0
	pairCount := 0
	length := len(substrings[0])

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := 0; k < length; k++ {
				if substrings[i][k] == substrings[j][k] {
					totalCommon++
				}
			}
			pairCount++
		}
	}

	totalPositions := pairCount * length
	percentage := (float64(totalCommon) / float64(totalPositions)) * 100

	return fmt.Sprintf("%.10f", percentage)
}