package kata

import (
	"math"
	"strconv"
)

func MaxSumDig(nMax, maxSum int) []int {
	var validNumbers []int

	// Find all valid numbers in range [1000, nMax]
	for num := 1000; num <= nMax; num++ {
		if isValid(num, maxSum) {
			validNumbers = append(validNumbers, num)
		}
	}

	if len(validNumbers) == 0 {
		return []int{0, 0, 0}
	}

	// Calculate total sum
	totalSum := 0
	for _, num := range validNumbers {
		totalSum += num
	}

	// Calculate mean
	mean := float64(totalSum) / float64(len(validNumbers))

	// Find closest to mean (if tie, pick smallest number)
	closest := validNumbers[0]
	minDiff := math.Abs(float64(closest) - mean)

	for _, num := range validNumbers {
		diff := math.Abs(float64(num) - mean)
		if diff < minDiff || (diff == minDiff && num < closest) {
			minDiff = diff
			closest = num
		}
	}

	return []int{len(validNumbers), closest, totalSum}
}

func isValid(num, maxSum int) bool {
	s := strconv.Itoa(num)
	digits := make([]int, len(s))

	for i, ch := range s {
		digits[i] = int(ch - '0')
	}

	// Check all windows of 4 consecutive digits
	for i := 0; i <= len(digits)-4; i++ {
		sum := digits[i] + digits[i+1] + digits[i+2] + digits[i+3]
		if sum > maxSum {
			return false
		}
	}

	return true
}