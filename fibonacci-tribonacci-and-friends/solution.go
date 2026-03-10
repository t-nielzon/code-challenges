package kata

func Xbonacci(signature []float64, n int) []float64 {
	x := len(signature)
	if n <= 0 {
		return []float64{}
	}
	result := make([]float64, n)
	copy(result, signature)
	for i := x; i < n; i++ {
		sum := 0.0
		for j := i - x; j < i; j++ {
			sum += result[j]
		}
		result[i] = sum
	}
	return result
}