package kata

// the sum of binomial coefficients along diagonal p starting at line p
// equals C(n+1, p+1) by the hockey stick identity
func Diagonal(n, p int) int {
	return binomial(n+1, p+1)
}

func binomial(n, k int) int {
	if k > n-k {
		k = n - k
	}
	result := 1
	for i := 0; i < k; i++ {
		result = result * (n - i) / (i + 1)
	}
	return result
}