package kata

import "math"

func DigPow(n, p int) int {
	sum := 0
	tmp := n
	digits := []int{}
	for tmp > 0 {
		digits = append([]int{tmp % 10}, digits...)
		tmp /= 10
	}
	for i, d := range digits {
		sum += int(math.Pow(float64(d), float64(p+i)))
	}
	if sum%n == 0 {
		return sum / n
	}
	return -1
}