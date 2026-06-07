package kata

import "errors"

func isBouncy(n int) bool {
	increasing, decreasing := true, true
	prev := n % 10
	n /= 10
	for n > 0 {
		d := n % 10
		if d > prev {
			decreasing = false
		}
		if d < prev {
			increasing = false
		}
		prev = d
		n /= 10
	}
	return !increasing && !decreasing
}

func bouncyRatio(ratio float64) (int, error) {
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