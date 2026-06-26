package kata

func Xbonacci(signature []float64, n int) []float64 {
	x := len(signature)
	if n <= 0 || x == 0 {
		return []float64{}
	}

	result := make([]float64, 0, n)
	for i := 0; i < x && i < n; i++ {
		result = append(result, signature[i])
	}

	for len(result) < n {
		var sum float64
		for i := len(result) - x; i < len(result); i++ {
			sum += result[i]
		}
		result = append(result, sum)
	}

	return result
}