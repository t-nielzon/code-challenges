package main

func FindPrimesSextuplet(sumLimit int) [6]int {
	isPrime := func(n int) bool {
		if n < 2 {
			return false
		}
		if n == 2 {
			return true
		}
		if n%2 == 0 {
			return false
		}
		if n == 3 {
			return true
		}
		if n%3 == 0 {
			return false
		}
		for i := 5; i*i <= n; i += 6 {
			if n%i == 0 || n%(i+2) == 0 {
				return false
			}
		}
		return true
	}

	p := 7
	for {
		if isPrime(p) && isPrime(p+4) && isPrime(p+6) && isPrime(p+10) && isPrime(p+12) && isPrime(p+16) {
			sum := 6*p + 48
			if sum > sumLimit {
				return [6]int{p, p + 4, p + 6, p + 10, p + 12, p + 16}
			}
		}
		p += 2
	}
}