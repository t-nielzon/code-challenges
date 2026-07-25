package main

func ColorChoice(m int64, n int64) int64 {
	binomial := func(n, k int64) int64 {
		if k > n || k < 0 {
			return 0
		}
		if k == 0 || k == n {
			return 1
		}
		if k > n-k {
			k = n - k
		}
		var result int64 = 1
		for i := int64(1); i <= k; i++ {
			result = result * (n - i + 1) / i
		}
		return result
	}

	for x := int64(0); x <= n; x++ {
		if binomial(n, x) == m {
			if x <= n-x {
				return x
			} else {
				return n - x
			}
		}
	}

	return -1
}