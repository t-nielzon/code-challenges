package kata

import "strconv"

func DigPow(n, p int) int {
	digits := strconv.Itoa(n)
	sum := 0
	exp := p
	for _, c := range digits {
		d := int(c - '0')
		term := 1
		for i := 0; i < exp; i++ {
			term *= d
		}
		sum += term
		exp++
	}
	if sum%n == 0 {
		return sum / n
	}
	return -1
}