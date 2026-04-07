package kata

func FaultyOdometer(n int) int {
	result := 0
	multiplier := 1
	for n > 0 {
		d := n % 10
		if d > 4 {
			d--
		}
		result += d * multiplier
		multiplier *= 9
		n /= 10
	}
	return result
}