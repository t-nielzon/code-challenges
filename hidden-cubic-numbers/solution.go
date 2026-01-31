package kata

import (
	"strconv"
	"strings"
	"unicode"
)

func IsSumOfCubes(s string) string {
	// Find all cubic numbers (0-999 where sum of cubes of digits equals number)
	isCubic := func(n int) bool {
		sum := 0
		temp := n
		for temp > 0 {
			d := temp % 10
			sum += d * d * d
			temp /= 10
		}
		// Special case: 0 is cubic (0^3 = 0)
		if n == 0 {
			return true
		}
		return sum == n
	}

	// Extract groups of digits from the string
	var groups []string
	var current strings.Builder

	for _, ch := range s {
		if unicode.IsDigit(ch) {
			current.WriteRune(ch)
			if current.Len() == 3 {
				groups = append(groups, current.String())
				current.Reset()
			}
		} else {
			if current.Len() > 0 {
				groups = append(groups, current.String())
				current.Reset()
			}
		}
	}
	if current.Len() > 0 {
		groups = append(groups, current.String())
	}

	// Find cubic numbers and calculate sum
	var cubicNumbers []int
	sum := 0

	for _, g := range groups {
		n, _ := strconv.Atoi(g)
		if isCubic(n) {
			cubicNumbers = append(cubicNumbers, n)
			sum += n
		}
	}

	if len(cubicNumbers) == 0 {
		return "Unlucky"
	}

	// Build result string
	var result strings.Builder
	for _, n := range cubicNumbers {
		result.WriteString(strconv.Itoa(n))
		result.WriteString(" ")
	}
	result.WriteString(strconv.Itoa(sum))
	result.WriteString(" Lucky")

	return result.String()
}