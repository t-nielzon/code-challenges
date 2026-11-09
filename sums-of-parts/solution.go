package kata

func PartsSums(ls []int) []int {
	result := make([]int, len(ls)+1)
	sum := 0
	
	// iterate from the end backwards to accumulate sums
	for i := len(ls) - 1; i >= 0; i-- {
		sum += ls[i]
		result[i] = sum
	}
	
	return result
}