package main

func Gap(g int, m int, n int) []int {
	isPrime := func(num int) bool {
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

	var prevPrime int
	found := false

	for i := m; i <= n; i++ {
		if isPrime(i) {
			if found && i-prevPrime == g {
				return []int{prevPrime, i}
			}
			prevPrime = i
			found = true
		}
	}

	return nil
}