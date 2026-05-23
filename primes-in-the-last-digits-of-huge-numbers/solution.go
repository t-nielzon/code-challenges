package kata

func KthLastDigPrime(k int) []int {
	const mod uint64 = 1000000000
	a, b, c, d, e := uint64(0), uint64(1), uint64(1), uint64(2), uint64(4)
	n := 4
	count := 0
	for {
		n++
		next := (e + d + b + 2*mod - c - a) % mod
		a, b, c, d, e = b, c, d, e, next
		if isPrime(next) {
			count++
			if count == k {
				return []int{n, int(next)}
			}
		}
	}
}

func mulMod(a, b, m uint64) uint64 {
	return (a * b) % m
}

func modPow(base, exp, m uint64) uint64 {
	result := uint64(1)
	base %= m
	for exp > 0 {
		if exp&1 == 1 {
			result = mulMod(result, base, m)
		}
		exp >>= 1
		base = mulMod(base, base, m)
	}
	return result
}

func isPrime(n uint64) bool {
	if n < 2 {
		return false
	}
	smallPrimes := []uint64{2, 3, 5, 7, 11, 13}
	for _, p := range smallPrimes {
		if n == p {
			return true
		}
		if n%p == 0 {
			return false
		}
	}
	d := n - 1
	s := 0
	for d%2 == 0 {
		d /= 2
		s++
	}
	witnesses := []uint64{2, 3, 5, 7, 11, 13}
	for _, w := range witnesses {
		if w >= n {
			continue
		}
		x := modPow(w, d, n)
		if x == 1 || x == n-1 {
			continue
		}
		composite := true
		for r := 0; r < s-1; r++ {
			x = mulMod(x, x, n)
			if x == n-1 {
				composite = false
				break
			}
		}
		if composite {
			return false
		}
	}
	return true
}