package kata

import "errors"

func isBouncy(n int) bool {
	increasing := true
	decreasing := true

	prev := n % 10
	n /= 10

	for n > 0 {
		digit := n % 10
		if digit > prev {
			increasing = false
		}
		if digit < prev {
			decreasing = false
		}
		if !increasing && !decreasing {
			return true
		}
		prev = digit
		n /= 10
	}

	return !increasing && !decreasing
}

func BouncyRatio(ratio float64) (int, error) {
	if ratio < 0 || ratio > 0.99 {
		return 0, errors.New("ratio must be between 0% and 99%")
	}

	bouncy := 0
	for n := 1; ; n++ {
		if isBouncy(n) {
			bouncy++
		}
		if float64(bouncy)/float64(n) >= ratio {
			return n, nil
		}
	}
}