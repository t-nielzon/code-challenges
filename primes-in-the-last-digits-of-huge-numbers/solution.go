package main

func KthLastDigPrime(k int) [2]int {
	modPow := func(base, exp, mod int64) int64 {
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

	isPrime := func(n int64) bool {
		if n < 2 {
			return false
		}
		if n == 2 || n == 3 {
			return true
		}
		if n%2 == 0 {
			return false
		}

		d := n - 1
		r := 0
		for d%2 == 0 {
			d /= 2
			r++
		}

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

	const MOD int64 = 1000000000

	f0, f1, f2, f3, f4 := int64(0), int64(1), int64(1), int64(2), int64(4)

	count := 0
	n := 5

	for {
		fn := (f4 + f3 - f2 + f1 - f0) % MOD
		if fn < 0 {
			fn += MOD
		}

		f0, f1, f2, f3, f4 = f1, f2, f3, f4, fn

		if isPrime(f4) {
			count++
			if count == k {
				return [2]int{n, int(f4)}
			}
		}

		n++
	}
}