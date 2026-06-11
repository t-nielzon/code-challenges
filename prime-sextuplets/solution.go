// solution.go

package kata

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n < 4 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func FindPrimesSextuplet(sumLimit int) []int {
	// the sextuplet sum is 6p + 48, which grows with p, so the first
	// sextuplet surpassing the limit is the one with the smallest valid p
	// whose sum exceeds sumLimit — start searching near that bound
	start := (sumLimit-48)/6 + 1
	if start < 7 {
		start = 7
	}
	// for p > 5, all six members avoid divisibility by 2, 3 and 5 only
	// when p ≡ 7 (mod 30), so candidates can step by 30
	p := start + ((7-start%30)+30)%30
	for ; ; p += 30 {
		if isPrime(p) && isPrime(p+4) && isPrime(p+6) &&
			isPrime(p+10) && isPrime(p+12) && isPrime(p+16) {
			return []int{p, p + 4, p + 6, p + 10, p + 12, p + 16}
		}
	}
}