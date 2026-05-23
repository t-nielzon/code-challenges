package kata

import (
	"fmt"
	"strconv"
	"strings"
)

func Balance(book string) string {
	cleaned := strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '.' || r == ' ' || r == '\n' {
			return r
		}
		return -1
	}, book)

	lines := strings.Split(cleaned, "\n")
	var nonBlank []string
	for _, l := range lines {
		trimmed := strings.TrimSpace(l)
		if trimmed != "" {
			nonBlank = append(nonBlank, trimmed)
		}
	}

	balance, _ := strconv.ParseFloat(strings.Fields(nonBlank[0])[0], 64)
	result := []string{fmt.Sprintf("Original Balance: %.2f", balance)}

	totalExpense := 0.0
	count := 0
	for _, line := range nonBlank[1:] {
		fields := strings.Fields(line)
		amount, _ := strconv.ParseFloat(fields[len(fields)-1], 64)
		balance -= amount
		totalExpense += amount
		count++
		result = append(result, fmt.Sprintf("%s Balance %.2f", strings.Join(fields, " "), balance))
	}

	result = append(result, fmt.Sprintf("Total expense  %.2f", totalExpense))
	avg := 0.0
	if count > 0 {
		avg = totalExpense / float64(count)
	}
	result = append(result, fmt.Sprintf("Average expense  %.2f", avg))

	return strings.Join(result, "\n")
}