package kata

func Diagonal(n, p int) int {
	// hockey stick identity: sum_{k=p}^{n} C(k, p) = C(n+1, p+1)
	result := 1
	for i := 0; i <= p; i++ {
		result = result * (n + 1 - i) / (i + 1)
	}
	return result
}