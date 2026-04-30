package kata

func Xbonacci(signature []int, n int) []int {
	if n <= 0 {
		return []int{}
	}
	x := len(signature)
	if n <= x {
		out := make([]int, n)
		copy(out, signature)
		return out
	}
	result := make([]int, n)
	copy(result, signature)
	for i := x; i < n; i++ {
		sum := 0
		for j := i - x; j < i; j++ {
			sum += result[j]
		}
		result[i] = sum
	}
	return result
}