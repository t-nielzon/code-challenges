package main

import (
	"regexp"
	"strconv"
	"strings"
)

func FindCubic(s string) string {
	// Extract all groups of consecutive digits
	re := regexp.MustCompile(`\d+`)
	digitGroups := re.FindAllString(s, -1)

	var cubicNumbers []int
	sum := 0

	// For each group of digits, split into chunks of at most 3
	for _, group := range digitGroups {
		// Split group into chunks of 3 from the left
		for i := 0; i < len(group); i += 3 {
			end := i + 3
			if end > len(group) {
				end = len(group)
			}
			chunk := group[i:end]
			num, _ := strconv.Atoi(chunk)

			// Check if it's cubic
			if isCubic(num) {
				cubicNumbers = append(cubicNumbers, num)
				sum += num
			}
		}
	}

	if len(cubicNumbers) == 0 {
		return "Unlucky"
	}

	// Build result
	var result []string
	for _, num := range cubicNumbers {
		result = append(result, strconv.Itoa(num))
	}
	result = append(result, strconv.Itoa(sum), "Lucky")

	return strings.Join(result, " ")
}

func isCubic(n int) bool {
	if n < 0 || n > 999 {
		return false
	}

	sum := 0
	temp := n

	for temp > 0 {
		digit := temp % 10
		sum += digit * digit * digit
		temp /= 10
	}

	return sum == n
}