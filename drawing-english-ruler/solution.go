package kata

import (
	"fmt"
	"strings"
)

func DrawRuler(t, n int) string {
	if t <= 0 {
		return ""
	}

	var lines []string

	// draw the interval between two major ticks recursively
	var drawInterval func(length int)
	drawInterval = func(length int) {
		if length < 1 {
			return
		}
		drawInterval(length - 1)
		lines = append(lines, strings.Repeat("-", length))
		drawInterval(length - 1)
	}

	for i := 0; i <= n; i++ {
		lines = append(lines, fmt.Sprintf("%s %d", strings.Repeat("-", t), i))
		if i < n {
			drawInterval(t - 1)
		}
	}

	return strings.Join(lines, "\n")
}