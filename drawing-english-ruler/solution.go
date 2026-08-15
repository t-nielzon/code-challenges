package kata

import (
	"strconv"
	"strings"
)

func DrawRuler(t, n int) string {
	var lines []string

	// Draw major tick 0
	lines = append(lines, strings.Repeat("-", t)+" 0")

	// For each interval between consecutive inches
	for i := 0; i < n; i++ {
		// Draw subdivisions between inch i and i+1
		addTicks(&lines, t-1)

		// Draw major tick i+1
		lines = append(lines, strings.Repeat("-", t)+" "+strconv.Itoa(i+1))
	}

	return strings.Join(lines, "\n")
}

func addTicks(lines *[]string, height int) {
	if height == 0 {
		return
	}

	if height == 1 {
		*lines = append(*lines, "-")
	} else {
		addTicks(lines, height-1)
		*lines = append(*lines, strings.Repeat("-", height))
		addTicks(lines, height-1)
	}
}