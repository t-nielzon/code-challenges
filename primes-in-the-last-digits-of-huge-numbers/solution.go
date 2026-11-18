package main

func kthLastDigPrime(k int) [2]int {
	const mod = 1_000_000_000

	f := [5]int64{0, 1, 1, 2, 4}

	count := 0
	n := 5

	for {
		// f(n) = f(n-1) + f(n-2) - f(n-3) + f(n-4) - f(n-5)
		next := f[0] + f[1] - f[2] + f[3] - f[4]

		// Keep in modulo range to get last 9 digits
		next = ((next % int64(mod)) + int64(mod)) % int64(mod)

		// Check if last 9 digits form a prime
		if isPrime(uint64(next)) {
			count++
			if count == k {
				return [2]int{n, int(next)}
			}
		}

		// Shift values
		f[4] = f[3]
		f[3] = f[2]
		f[2] = f[1]
		f[1] = f[0]
		f[0] = next

		n++
	}
}

func isPrime(n uint64) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	// Write n-1 as 2^r * d
	r := 0
	d := n - 1
	for d%2 == 0 {
		d /= 2
		r++
	}

	// Miller-Rabin test with deterministic witnesses for n < 3,317,044,064,679,887,385,961,981
	witnesses := []uint64{2, 3, 5, 7, 11, 13, 17}

	for _, a := range witnesses {
		if a >= n {
			continue
		}

		x := modExp(a, d, n)

		if x == 1 || x == n-1 {
			continue
		}

		composite := true
		for i := 0; i < r-1; i++ {
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

func modExp(base, exp, mod uint64) uint64 {
	result := uint64(1)
	base = base % mod
	for exp > 0 {
		if exp%2 == 1 {
			result = mulMod(result, base, mod)
		}
		exp >>= 1
		base = mulMod(base, base, mod)
	}
	return result
}

func mulMod(a, b, mod uint64) uint64 {
	return (a % mod) * (b % mod) % mod
}