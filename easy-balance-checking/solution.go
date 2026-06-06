package kata

import (
	"fmt"
	"strconv"
	"strings"
)

func Balance(book string) string {
	// keep only letters, digits, dots and spaces
	cleaned := strings.Map(func(r rune) rune {
		switch {
		case r >= 'a' && r <= 'z',
			r >= 'A' && r <= 'Z',
			r >= '0' && r <= '9',
			r == '.', r == ' ', r == '\n':
			return r
		default:
			return -1
		}
	}, book)

	lines := strings.Split(cleaned, "\n")

	var report []string
	balance := 0.0
	total := 0.0
	count := 0

	for i, line := range lines {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}

		if i == 0 {
			balance, _ = strconv.ParseFloat(fields[0], 64)
			report = append(report, fmt.Sprintf("Original Balance: %.2f", balance))
			continue
		}

		// last field is the amount, preceding fields are number + category
		amount, _ := strconv.ParseFloat(fields[len(fields)-1], 64)
		balance -= amount
		total += amount
		count++

		report = append(report, fmt.Sprintf("%s Balance %.2f", strings.Join(fields, " "), balance))
	}

	average := 0.0
	if count > 0 {
		average = total / float64(count)
	}

	report = append(report, fmt.Sprintf("Total expense  %.2f", total))
	report = append(report, fmt.Sprintf("Average expense  %.2f", average))

	return strings.Join(report, "\n")
}