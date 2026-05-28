package kata

func Tribonacci(signature [3]int, n int) []int {
	result := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if i < 3 {
			result = append(result, signature[i])
		} else {
			result = append(result, result[i-1]+result[i-2]+result[i-3])
		}
	}
	return result
}