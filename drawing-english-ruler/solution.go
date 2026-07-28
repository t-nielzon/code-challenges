package main

import (
	"fmt"
	"strings"
)

func DrawRuler(t int, n int) string {
	var lines []string

	// Draw the major tick at position 0
	lines = append(lines, strings.Repeat("-", t)+" 0")

	// For each inch
	for i := 0; i < n; i++ {
		// Draw the sub-ticks in this interval
		for _, tickLength := range ticksInRange(0, 1, t-1) {
			lines = append(lines, strings.Repeat("-", tickLength))
		}

		// Draw the major tick for the next inch
		lines = append(lines, fmt.Sprintf("%s %d", strings.Repeat("-", t), i+1))
	}

	return strings.Join(lines, "\n")
}

func ticksInRange(startPos, endPos float64, tickLength int) []int {
	if tickLength == 0 {
		return []int{}
	}

	midPos := (startPos + endPos) / 2
	result := ticksInRange(startPos, midPos, tickLength-1)
	result = append(result, tickLength)
	result = append(result, ticksInRange(midPos, endPos, tickLength-1)...)

	return result
}