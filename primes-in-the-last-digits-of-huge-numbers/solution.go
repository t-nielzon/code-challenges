package main

func modPow(base, exp, mod int64) int64 {
	result := int64(1)
	base = base % mod
	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % mod
		}
		exp = exp >> 1
		base = (base * base) % mod
	}
	return result
}

func isPrime(n int64) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	// write n-1 as 2^r * d
	r := 0
	d := n - 1
	for d%2 == 0 {
		r++
		d /= 2
	}

	// test against these witnesses for deterministic result up to 2^64
	witnesses := []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}

	for _, a := range witnesses {
		if a >= n {
			continue
		}

		x := modPow(a, d, n)
		if x == 1 || x == n-1 {
			continue
		}

		composite := true
		for i := 0; i < r-1; i++ {
			x = (x * x) % n
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

func kthLastDigPrime(k int) []int64 {
	const MOD = int64(1000000000) // 10^9

	f := [5]int64{0, 1, 1, 2, 4}
	count := 0
	n := 5

	for {
		// f(n) = f(n-1) + f(n-2) - f(n-3) + f(n-4) - f(n-5)
		next := (f[4] + f[3] - f[2] + f[1] - f[0]) % MOD
		if next < 0 {
			next += MOD
		}

		if isPrime(next) {
			count++
			if count == k {
				return []int64{int64(n), next}
			}
		}

		f[0] = f[1]
		f[1] = f[2]
		f[2] = f[3]
		f[3] = f[4]
		f[4] = next

		n++
	}
}