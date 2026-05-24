package kata

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	if x < 4 {
		return true
	}
	if x%2 == 0 {
		return false
	}
	for i := 3; i*i <= x; i += 2 {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func Step(g, m, n int) []int {
	for i := m; i+g <= n; i++ {
		if isPrime(i) && isPrime(i+g) {
			return []int{i, i + g}
		}
	}
	return nil
}