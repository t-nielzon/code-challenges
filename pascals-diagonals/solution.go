package kata

func Generate(n, l int) []uint64 {
	result := make([]uint64, 0, l)
	if l == 0 {
		return result
	}
	for i := 0; i < l; i++ {
		result = append(result, binomial(n+i, n))
	}
	return result
}

func binomial(n, k int) uint64 {
	if k > n-k {
		k = n - k
	}
	var c uint64 = 1
	for i := 0; i < k; i++ {
		c = c * uint64(n-i) / uint64(i+1)
	}
	return c
}