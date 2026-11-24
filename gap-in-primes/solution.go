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

	// find the first prime >= m
	for i := m; i <= n; i++ {
		if isPrime(i) {
			prevPrime = i
			break
		}
	}

	// if no first prime found, return nil
	if prevPrime == 0 {
		return nil
	}

	// look for the next prime with the desired gap
	for i := prevPrime + 1; i <= n; i++ {
		if isPrime(i) {
			if i-prevPrime == g {
				return []int{prevPrime, i}
			}
			prevPrime = i
		}
	}

	return nil
}