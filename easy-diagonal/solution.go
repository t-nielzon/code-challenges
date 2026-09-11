func diagonal(n, p int) int {
	// apply hockey stick identity: sum equals binomial coefficient
	return binomialCoeff(n+1, p+1)
}

func binomialCoeff(n, k int) int {
	if k > n {
		return 0
	}
	if k == 0 || k == n {
		return 1
	}
	
	if k > n-k {
		k = n - k
	}
	
	result := 1
	for i := 0; i < k; i++ {
		result = result * (n - i) / (i + 1)
	}
	return result
}