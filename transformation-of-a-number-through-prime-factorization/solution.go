package kata

func F(n int) int {
	result := 1
	hasExpGreaterThanOne := false
	num := n
	for p := 2; p*p <= num; p++ {
		if num%p == 0 {
			k := 0
			for num%p == 0 {
				num /= p
				k++
			}
			if k > 1 {
				hasExpGreaterThanOne = true
				factor := k
				for i := 0; i < k-1; i++ {
					factor *= p
				}
				result *= factor
			}
		}
	}
	if num > 1 {
	}
	if !hasExpGreaterThanOne {
		return 1
	}
	return result
}