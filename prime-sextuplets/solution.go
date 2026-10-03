package kata

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func FindPrimesSextuplet(sumLimit int) [6]int {
	p := 7
	for {
		// all sextuplets have p ≡ 2 (mod 5) due to divisibility by 5 in other positions
		if p%5 == 2 && isPrime(p) && isPrime(p+4) && isPrime(p+6) && isPrime(p+10) && isPrime(p+12) && isPrime(p+16) {
			sum := 6*p + 48
			if sum > sumLimit {
				return [6]int{p, p+4, p+6, p+10, p+12, p+16}
			}
		}
		p += 2
	}
}