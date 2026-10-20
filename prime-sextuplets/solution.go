package main

func FindPrimesSextuplet(sumLimit int) []int {
	offsets := []int{0, 4, 6, 10, 12, 16}
	
	for p := 7; ; p += 6 {
		allPrime := true
		sum := 0
		for _, offset := range offsets {
			candidate := p + offset
			if !isPrime(candidate) {
				allPrime = false
				break
			}
			sum += candidate
		}
		
		if allPrime && sum > sumLimit {
			return []int{p, p + 4, p + 6, p + 10, p + 12, p + 16}
		}
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 {
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