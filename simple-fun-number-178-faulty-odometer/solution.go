// solution.go
package kata

func FaultyOdometer(n int) int {
	result := 0
	base := 1
	for n > 0 {
		d := n % 10
		if d > 4 {
			d--
		}
		result += d * base
		base *= 9
		n /= 10
	}
	return result
}