package kata

func isPrime(n uint64) bool {
	if n < 2 {
		return false
	}
	if n < 4 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := uint64(3); i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func reverseNum(n uint64) uint64 {
	r := uint64(0)
	for n > 0 {
		r = r*10 + n%10
		n /= 10
	}
	return r
}

func BackwardsPrime(start, stop uint64) []uint64 {
	result := []uint64{}
	if start < 2 {
		start = 2
	}
	for n := start; n <= stop; n++ {
		if n < 10 {
			continue
		}
		r := reverseNum(n)
		if r == n {
			continue
		}
		if isPrime(n) && isPrime(r) {
			result = append(result, n)
		}
	}
	return result
}