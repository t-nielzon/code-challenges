package kata

import (
	"fmt"
	"strings"
)

func Histogram(results []int) string {
	// count occurrences for each die face
	counts := make([]int, 6)
	for _, r := range results {
		if r >= 1 && r <= 6 {
			counts[r-1]++
		}
	}

	max := 0
	for _, c := range counts {
		if c > max {
			max = c
		}
	}

	if max == 0 {
		return "-----------\n1 2 3 4 5 6\n"
	}

	// each column is 2 chars wide (value + space), except last column has no trailing space
	// columns are at positions 0, 2, 4, 6, 8, 10 (character indices)
	// label row goes one row above the bar, so at row (max - count)
	// bar rows go from (max - count + 1) to max

	var lines []string

	for row := 0; row <= max; row++ {
		// row 0 is the top, row max is the bottom bar row
		// for each column, figure out what goes there
		var cells [6]string
		for col := 0; col < 6; col++ {
			c := counts[col]
			if c == 0 {
				cells[col] = " "
				continue
			}
			// the label row for this column is at row = max - c
			// bar rows are max - c + 1 through max
			labelRow := max - c
			if row == labelRow {
				cells[col] = fmt.Sprintf("%d", c)
			} else if row > labelRow {
				cells[col] = "#"
			} else {
				cells[col] = " "
			}
		}

		// build line: each cell gets 2-char width except we trim trailing spaces
		// handle multi-char labels (counts can be >= 10)
		// use a fixed-width approach: each column occupies 2 characters
		// but labels like "10" occupy 2 chars themselves
		var buf strings.Builder
		for col := 0; col < 6; col++ {
			if col > 0 {
				// space separator between columns, but label might overflow
				// the column positions: col 0 starts at 0, col 1 at 2, etc.
				// we need to handle the case where a label is wider than 1 char
			}
			s := cells[col]
			if col == 0 {
				buf.WriteString(s)
			} else {
				// pad so this column starts at position col*2
				for buf.Len() < col*2 {
					buf.WriteByte(' ')
				}
				buf.WriteString(s)
			}
		}

		line := strings.TrimRight(buf.String(), " ")
		lines = append(lines, line)
	}

	// add separator and labels
	lines = append(lines, "-----------")
	lines = append(lines, "1 2 3 4 5 6")

	// remove leading empty lines
	start := 0
	for start < len(lines) && lines[start] == "" {
		start++
	}

	var result strings.Builder
	for i := start; i < len(lines); i++ {
		result.WriteString(lines[i])
		result.WriteByte('\n')
	}

	return result.String()
}