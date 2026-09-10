package main

func Step(g, m, n int) []int {
	for i := m; i <= n-g; i++ {
		if isPrime(i) && isPrime(i+g) {
			return []int{i, i + g}
		}
	}
	return nil
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}
	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}