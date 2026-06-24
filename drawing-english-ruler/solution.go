package kata

import (
	"strconv"
	"strings"
)

func DrawRuler(t, n int) string {
	var lines []string

	drawLine := func(tickLength int, label string) {
		dashes := strings.Repeat("-", tickLength)
		if label != "" {
			lines = append(lines, dashes+" "+label)
		} else {
			lines = append(lines, dashes)
		}
	}

	var drawInterval func(centerLength int)
	drawInterval = func(centerLength int) {
		if centerLength > 0 {
			drawInterval(centerLength - 1)
			drawLine(centerLength, "")
			drawInterval(centerLength - 1)
		}
	}

	drawLine(t, "0")
	for j := 1; j <= n; j++ {
		drawInterval(t - 1)
		drawLine(t, strconv.Itoa(j))
	}

	return strings.Join(lines, "\n")
}