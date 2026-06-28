package kata

func Step(g, m, n int) []int {
	isPrime := func(num int) bool {
		if num < 2 {
			return false
		}
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}

	for i := m; i <= n-g; i++ {
		if isPrime(i) && isPrime(i+g) {
			return []int{i, i + g}
		}
	}
	return nil
}