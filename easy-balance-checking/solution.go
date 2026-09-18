package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Balance(s string) string {
	lines := strings.Split(s, "\n")
	
	// Parse initial balance
	initialBalance, _ := strconv.ParseFloat(strings.TrimSpace(lines[0]), 64)
	currentBalance := initialBalance
	
	var report []string
	report = append(report, fmt.Sprintf("Original Balance: %.2f", initialBalance))
	
	var expenses []float64
	
	// Process each check
	for i := 1; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}
		
		// Clean the line - keep only alphanumeric, dots, and spaces
		re := regexp.MustCompile(`[^a-zA-Z0-9. ]`)
		cleanedLine := re.ReplaceAllString(line, "")
		cleanedLine = strings.TrimSpace(cleanedLine)
		
		// Parse the cleaned line
		parts := strings.Fields(cleanedLine)
		if len(parts) < 3 {
			continue
		}
		
		checkNum := parts[0]
		amount, _ := strconv.ParseFloat(parts[len(parts)-1], 64)
		category := strings.Join(parts[1:len(parts)-1], " ")
		
		currentBalance -= amount
		expenses = append(expenses, amount)
		
		report = append(report, fmt.Sprintf("%s %s %.2f Balance %.2f", checkNum, category, amount, currentBalance))
	}
	
	// Calculate total and average
	total := 0.0
	for _, exp := range expenses {
		total += exp
	}
	
	avg := 0.0
	if len(expenses) > 0 {
		avg = total / float64(len(expenses))
	}
	
	report = append(report, fmt.Sprintf("Total expense  %.2f", total))
	report = append(report, fmt.Sprintf("Average expense  %.2f", avg))
	
	return strings.Join(report, "\n")
}