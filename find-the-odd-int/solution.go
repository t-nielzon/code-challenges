package kata

func FindOdd(arr []int) int {
	result := 0
	for _, num := range arr {
		result ^= num
	}
	return result
}