package kata

import (
	"strconv"
	"strings"
)

func DrawRuler(t, n int) string {
	if t <= 0 || n < 0 {
		return ""
	}

	var lines []string
	lines = append(lines, strings.Repeat("-", t)+" "+strconv.Itoa(0))

	for i := 1; i <= n; i++ {
		drawMiddle(&lines, t-1)
		lines = append(lines, strings.Repeat("-", t)+" "+strconv.Itoa(i))
	}

	return strings.Join(lines, "\n")
}

func drawMiddle(lines *[]string, length int) {
	if length <= 0 {
		return
	}
	drawMiddle(lines, length-1)
	*lines = append(*lines, strings.Repeat("-", length))
	drawMiddle(lines, length-1)
}