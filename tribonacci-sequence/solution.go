package kata

func Tribonacci(signature [3]float64, n int) []float64 {
	if n == 0 {
		return []float64{}
	}
	result := make([]float64, n)
	for i := 0; i < n && i < 3; i++ {
		result[i] = signature[i]
	}
	for i := 3; i < n; i++ {
		result[i] = result[i-1] + result[i-2] + result[i-3]
	}
	return result
}