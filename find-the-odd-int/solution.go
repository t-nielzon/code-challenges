package kata

func FindOdd(seq []int) int {
	result := 0
	for _, n := range seq {
		result ^= n
	}
	return result
}