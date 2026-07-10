package main

import (
	"fmt"
	"strings"
)

type Tick struct {
	Length int
	Label  *int
}

func DrawRuler(t, n int) string {
	var ticks []Tick
	
	// Add major ticks
	for i := 0; i <= n; i++ {
		label := i
		ticks = append(ticks, Tick{Length: t, Label: &label})
	}
	
	// Add subdivision ticks for each interval
	for i := 0; i < n; i++ {
		addSubdivisions(t-1, &ticks)
	}
	
	// Format output
	var lines []string
	for _, tick := range ticks {
		line := strings.Repeat("-", tick.Length)
		if tick.Label != nil {
			line += " " + fmt.Sprint(*tick.Label)
		}
		lines = append(lines, line)
	}
	
	return strings.Join(lines, "\n")
}

func addSubdivisions(length int, ticks *[]Tick) {
	if length <= 0 {
		return
	}
	
	addSubdivisions(length-1, ticks)
	*ticks = append(*ticks, Tick{Length: length})
	addSubdivisions(length-1, ticks)
}