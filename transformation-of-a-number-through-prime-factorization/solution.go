package kata

func F(n int) int {
	result := 1
	x := n
	for p := 2; p*p <= x; p++ {
		if x%p == 0 {
			k := 0
			for x%p == 0 {
				x /= p
				k++
			}
			factor := k
			for i := 0; i < k-1; i++ {
				factor *= p
			}
			result *= factor
		}
	}
	if x > 1 {
		result *= 1
	}
	return result
}