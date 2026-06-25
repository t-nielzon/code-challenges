package kata

import "math"

func DigPow(n, p int) int {
	digits := []int{}
	for x := n; x > 0; x /= 10 {
		digits = append([]int{x % 10}, digits...)
	}

	sum := 0
	for i, d := range digits {
		sum += int(math.Pow(float64(d), float64(p+i)))
	}

	if sum%n == 0 {
		return sum / n
	}
	return -1
}